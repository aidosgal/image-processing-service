package app

import (
	"log/slog"

	grpcapp "github.com/aidosgal/image-processing-service/internal/app/grpc"
	"github.com/aidosgal/image-processing-service/internal/config"
	"github.com/aidosgal/image-processing-service/internal/repository/psql"
	service "github.com/aidosgal/image-processing-service/internal/service/image"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, cfg config.DatabaseConfig) *App {
	reposiry, err := psql.NewRepository(cfg)
	if err != nil {
		panic(err)
	}

	service := service.NewImageService(log, reposiry)

	grpcApp := grpcapp.NewApp(log, service, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
