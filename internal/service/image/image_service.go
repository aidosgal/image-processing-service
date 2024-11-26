package service

import (
	"context"
	"log/slog"
)

type ImageService struct {
	log        *slog.Logger
	repository ImageRepository
}

type ImageRepository interface {
}

func NewImageService(log *slog.Logger, repository ImageRepository) *ImageService {
	return &ImageService{
		log:        log,
		repository: repository,
	}
}

func (i *ImageService) UploadImage(ctx context.Context, image []byte, filename string) (int64, error) {

	return 0, nil
}

func (i *ImageService) ListImages(ctx context.Context, image []byte, filename string) (int64, error) {

	return 0, nil
}

func (i *ImageService) GetImage(ctx context.Context, image []byte, filename string) (int64, error) {

	return 0, nil
}

func (i *ImageService) DeleteImage(ctx context.Context, image []byte, filename string) (int64, error) {

	return 0, nil
}
