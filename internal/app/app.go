package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	client "github.com/Makovey/go-keeper/internal/client/grpc"
	"github.com/Makovey/go-keeper/internal/client/ui"
	"github.com/Makovey/go-keeper/internal/config"
	grpcAuth "github.com/Makovey/go-keeper/internal/gen/auth"
	grpcStorage "github.com/Makovey/go-keeper/internal/gen/storage"
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

func (a *App) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	wg.Add(1)
	go a.runGRPCServer(ctx, &wg)

	wg.Add(1)
	go a.runUI(ctx, &wg)

	wg.Wait()
}

func (a *App) runGRPCServer(ctx context.Context, wg *sync.WaitGroup) {
	fn := "app.runGRPCServer"
	defer wg.Done()

	listen, err := net.Listen("tcp", a.cfg.GRPCPort())
	if err != nil {
		a.log.Errorf("[%s]: failed to listen: %s", fn, err.Error())
		return
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
	grpcAuth.RegisterAuthServer(s, a.authServer)
	grpcStorage.RegisterStorageServiceServer(s, a.storage)

	a.log.Infof("[%s]: starting grpc server on: %s", fn, a.cfg.GRPCPort())
	go func() {
		if err = s.Serve(listen); err != nil {
			a.log.Errorf("[%s]: can't serve grpc server: %s", fn, err.Error())
			return
		}
	}()

	<-ctx.Done()

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

func (a *App) runUI(ctx context.Context, wg *sync.WaitGroup) {
	fn := "app.runUI"
	defer wg.Done()

	conn, err := grpc.NewClient(
		a.cfg.ClientConnectionHost()+a.cfg.GRPCPort(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		a.log.Errorf("[%s]: failed to create grpc client: %s", fn, err.Error())
		return
	}
	defer conn.Close()

	authClient := client.NewAuthClient(a.log, grpcAuth.NewAuthClient(conn))
	storageClient := client.NewStorageClient(a.log, utils.NewDirManager(), grpcStorage.NewStorageServiceClient(conn))

	p := tea.NewProgram(ui.InitialModel(authClient, storageClient, a.cfg.UpdateDurationForUI()), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		a.log.Infof("[%s]: can't run ui program, cause: %v", fn, err)
		return
	}

	<-ctx.Done()
	p.Quit()
}
