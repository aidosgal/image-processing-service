package service

import (
	"context"
	"log/slog"

	imagev1 "github.com/aidosgal/image-processing-service/pkg/gen/go/image"
)

type ImageService struct {
	log        *slog.Logger
	repository Repository
}

type Repository interface {
}

func NewImageService(log *slog.Logger, repository Repository) *ImageService {
	return &ImageService{
		log:        log,
		repository: repository,
	}
}

func (i *ImageService) UploadImage(ctx context.Context, image []byte, filename string) (int64, error) {

	return 0, nil
}

func (i *ImageService) ListImages(ctx context.Context) ([]*imagev1.ImageMetadata, error) {

	return nil, nil
}

func (i *ImageService) GetImage(ctx context.Context, image_id int64) ([]byte, *imagev1.ImageMetadata, error) {

	return nil, nil, nil
}

func (i *ImageService) DeleteImage(ctx context.Context, image_id int64) (bool, error) {

	return false, nil
}
