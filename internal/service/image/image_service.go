package service

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	lib "github.com/aidosgal/image-processing-service/internal/lib/service"
	imagev1 "github.com/aidosgal/image-processing-service/pkg/gen/go/image"
)

type ImageService struct {
	log        *slog.Logger
	repository Repository
}

type Repository interface {
	StoreImage(ctx context.Context, metadata *imagev1.ImageMetadata) (int64, error)
	GetAllImages(ctx context.Context) ([]*imagev1.ImageMetadata, error)
	GetImageById(ctx context.Context, image_id int64) (*imagev1.ImageMetadata, error)
	DeleteImageById(ctx context.Context, image_id int64) (bool, error)
}

func NewImageService(log *slog.Logger, repository Repository) *ImageService {
	return &ImageService{
		log:        log,
		repository: repository,
	}
}

func (i *ImageService) UploadImage(ctx context.Context, image []byte, filename string) (int64, error) {
	uploadsDir := "./uploads/images"

	uniqueFilename := lib.GenerateUniqueFilename(filename)
	filePath := filepath.Join(uploadsDir, uniqueFilename)

	if err := os.WriteFile(filePath, image, 0644); err != nil {
		return 0, fmt.Errorf("failed to save image: %w", err)
	}

	var wg sync.WaitGroup
	var metadataErr error
	var thumbnailErr error
	var metadata *imagev1.ImageMetadata
	var thumbnailPath string

	metadataChan := make(chan *imagev1.ImageMetadata, 1)
	thumbnailChan := make(chan string, 1)
	errChan := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(metadataChan)

		extractedMetadata, err := lib.ExtractImageMetadata(filePath, uniqueFilename)
		if err != nil {
			metadataErr = fmt.Errorf("metadata extraction failed: %w", err)
			errChan <- metadataErr
			return
		}
		metadataChan <- extractedMetadata
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(thumbnailChan)

		generatedThumbnailPath, err := lib.GenerateThumbnail(filePath)
		if err != nil {
			thumbnailErr = fmt.Errorf("thumbnail generation failed: %w", err)
			errChan <- thumbnailErr
			return
		}
		thumbnailChan <- generatedThumbnailPath
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			os.Remove(filePath)
			return 0, err
		}
	}

	metadata = <-metadataChan
	thumbnailPath = <-thumbnailChan

	if thumbnailPath != "" {
		metadata.ThumbnailPath = thumbnailPath
	}

	return i.repository.StoreImage(ctx, metadata)
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
