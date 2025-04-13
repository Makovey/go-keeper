package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	client "github.com/Makovey/go-keeper/internal/client/grpc"
	"github.com/Makovey/go-keeper/internal/client/ui"
	"github.com/Makovey/go-keeper/internal/config"
	pbAuth "github.com/Makovey/go-keeper/internal/gen/auth"
	pbStorage "github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/interceptor/stream"
	"github.com/Makovey/go-keeper/internal/interceptor/unary"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/service/jwt"
	"github.com/Makovey/go-keeper/internal/transport/grpc/auth"
	"github.com/Makovey/go-keeper/internal/transport/grpc/storage"
	"github.com/Makovey/go-keeper/internal/utils"
)

type App struct {
	cfg        config.Config
	log        logger.Logger
	authServer *auth.Server
	storage    *storage.Server
}

func NewApp(
	cfg config.Config,
	log logger.Logger,
	authServer *auth.Server,
	storage *storage.Server,
) *App {
	return &App{
		cfg:        cfg,
		log:        log,
		authServer: authServer,
		storage:    storage,
	}
}

func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
	defer stop()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return a.runGRPCServer(ctx)
	})

	g.Go(func() error {
		return a.runUI(ctx)
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPCServer(ctx context.Context) error {
	fn := "app.runGRPCServer"

	listen, err := net.Listen("tcp", a.cfg.GRPCPort())
	if err != nil {
		a.log.Errorf("[%s]: failed to listen: %s", fn, err.Error())
		return err
	}

	s := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			unary.Logger(a.log),
			unary.JWTAuth(a.log, jwt.NewManager(a.cfg)),
		),
		grpc.ChainStreamInterceptor(
			stream.Logger(a.log),
			stream.JWTAuth(a.log, jwt.NewManager(a.cfg)),
		),
	)

	reflection.Register(s)
	pbAuth.RegisterAuthServer(s, a.authServer)
	pbStorage.RegisterStorageServiceServer(s, a.storage)

	a.log.Infof("[%s]: starting grpc server on: %s", fn, a.cfg.GRPCPort())
	serveErr := make(chan error, 1)
	go func() {
		if err = s.Serve(listen); err != nil {
			serveErr <- fmt.Errorf("[%s]: can't serve grpc server: %w", fn, err)
		}
	}()

	select {
	case err = <-serveErr:
		return err
	case <-ctx.Done():
		shutDownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		isStoppedCh := make(chan struct{})
		go func() {
			s.GracefulStop()
			close(isStoppedCh)
		}()

		select {
		case <-isStoppedCh:
			a.log.Infof("[%s]: grpc server stopped gracefully", fn)
		case <-shutDownCtx.Done():
			a.log.Errorf("[%s]: graceful shutdown timeout reached, forcing shutdown", fn)
			s.Stop()
		}
	}
	return nil
}

func (a *App) runUI(ctx context.Context) error {
	fn := "app.runUI"

	conn, err := grpc.NewClient(
		a.cfg.ClientConnectionHost()+a.cfg.GRPCPort(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		a.log.Errorf("[%s]: failed to create grpc client: %s", fn, err.Error())
		return err
	}
	defer conn.Close()

	authClient := client.NewAuthClient(a.log, pbAuth.NewAuthClient(conn))
	storageClient := client.NewStorageClient(a.log, utils.NewDirManager(), pbStorage.NewStorageServiceClient(conn))

	p := tea.NewProgram(ui.InitialModel(authClient, storageClient, a.cfg.UpdateDurationForUI()), tea.WithAltScreen(), tea.WithContext(ctx))
	if _, err = p.Run(); err != nil {
		a.log.Infof("[%s]: can't run ui program, cause: %v", fn, err)
		return err
	}

	<-ctx.Done()
	p.Quit()

	return nil
}
