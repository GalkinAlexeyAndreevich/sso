package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	authgrpc "sso/internal/grpc/auth"

	"google.golang.org/grpc"
)

type App struct {
	log  *slog.Logger
	gRPC *grpc.Server
	port int
}

func New(log *slog.Logger, authService authgrpc.Auth, port int) *App {
	grpcServer := grpc.NewServer()

	authgrpc.Register(grpcServer, authService)

	return &App{
		log:  log,
		gRPC: grpcServer,
		port: port,
	}
}

func (app *App) MustRun() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func (app *App) Run() error {
	const op = "grpcapp.Run"

	log := app.log.With(
		slog.String("op", op),
		slog.Int("port", app.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err := app.gRPC.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (app *App) Stop() {
	const op = "grpcapp.Stop"

	app.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", app.port))

	app.gRPC.GracefulStop()
}
