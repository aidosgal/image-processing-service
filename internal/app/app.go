package app

import (
	"log/slog"

	grpcapp "github.com/aidosgal/image-processing-service/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int) *App {
	grpcApp := grpcapp.NewApp(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
