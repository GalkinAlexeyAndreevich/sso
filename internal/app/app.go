package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: инициализировать хранилище (storage)
	// Пока оставил заглушку в виде nil, но надо добавить сервисный слой
	grpcApp := grpcapp.New(log, nil, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
