package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/Makovey/go-keeper/internal/config"
	grpc_auth "github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/transport/grpc/auth"
)

type App struct {
	cfg        config.Config
	log        logger.Logger
	authServer *auth.Server
}

func NewApp(
	cfg config.Config,
	log logger.Logger,
	authServer *auth.Server,
) *App {
	return &App{
		cfg:        cfg,
		log:        log,
		authServer: authServer,
	}
}

func (a *App) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	wg.Add(1)
	go a.runGRPCServer(ctx, &wg)

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
	)

	reflection.Register(s)
	grpc_auth.RegisterAuthServer(s, a.authServer)

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
