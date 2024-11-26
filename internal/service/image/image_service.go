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
	i.log.Info("Listing all images")

	images, err := i.repository.GetAllImages(ctx)
	if err != nil {
		i.log.Error("Failed to list images", "error", err)
		return nil, fmt.Errorf("failed to list images: %w", err)
	}

	i.log.Info("Images retrieved successfully", "count", len(images))

	return images, nil
}

func (i *ImageService) GetImage(ctx context.Context, imageID int64) ([]byte, *imagev1.ImageMetadata, error) {
	i.log.Info("Retrieving image", "image_id", imageID)

	metadata, err := i.repository.GetImageById(ctx, imageID)
	if err != nil {
		i.log.Error("Failed to retrieve image metadata", "image_id", imageID, "error", err)
		return nil, nil, fmt.Errorf("failed to retrieve image metadata: %w", err)
	}

	imageBytes, err := os.ReadFile(metadata.GetFilePath())
	if err != nil {
		i.log.Error("Failed to read image file", "image_path", metadata.GetFilePath(), "error", err)
		return nil, nil, fmt.Errorf("failed to read image file: %w", err)
	}

	i.log.Info("Image retrieved successfully", "image_id", imageID, "filename", metadata.GetFilename())

	return imageBytes, metadata, nil
}

func (i *ImageService) DeleteImage(ctx context.Context, imageID int64) (bool, error) {
	i.log.Info("Deleting image", "image_id", imageID)

	metadata, err := i.repository.GetImageById(ctx, imageID)
	if err != nil {
		i.log.Error("Failed to retrieve image metadata for deletion", "image_id", imageID, "error", err)
		return false, fmt.Errorf("failed to retrieve image metadata: %w", err)
	}

	var wg sync.WaitGroup
	var primaryFileErr, thumbnailErr error

	wg.Add(1)
	go func() {
		defer wg.Done()
		primaryFileErr = os.Remove(metadata.GetFilePath())
		if primaryFileErr != nil {
			i.log.Error("Failed to delete primary image file",
				"image_path", metadata.GetFilePath(),
				"error", primaryFileErr)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if metadata.GetThumbnailPath() != "" {
			thumbnailErr = os.Remove(metadata.GetThumbnailPath())
			if thumbnailErr != nil {
				i.log.Error("Failed to delete thumbnail",
					"thumbnail_path", metadata.GetThumbnailPath(),
					"error", thumbnailErr)
			}
		}
	}()

	wg.Wait()

	if primaryFileErr != nil || thumbnailErr != nil {
		i.log.Warn("Some files could not be deleted",
			"primary_file_error", primaryFileErr,
			"thumbnail_error", thumbnailErr)
	}

	deleted, err := i.repository.DeleteImageById(ctx, imageID)
	if err != nil {
		i.log.Error("Failed to delete image from database", "image_id", imageID, "error", err)
		return false, fmt.Errorf("failed to delete image from database: %w", err)
	}

	if deleted {
		i.log.Info("Image deleted successfully", "image_id", imageID)
	} else {
		i.log.Warn("Image not found or already deleted", "image_id", imageID)
	}

	return deleted, nil
}
