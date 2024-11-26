package lib

import (
	"context"
	"fmt"
	"image"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	imagev1 "github.com/aidosgal/image-processing-service/pkg/gen/go/image"
	"github.com/disintegration/imaging"
)

func GenerateUniqueFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	baseName := strings.TrimSuffix(originalFilename, ext)
	timestamp := time.Now().Format("20060102_150405")
	return fmt.Sprintf("%s_%s%s", baseName, timestamp, ext)
}

func getMimeType(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return "application/octet-stream"
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "application/octet-stream"
	}

	return http.DetectContentType(buffer)
}

func ExtractImageMetadata(filePath string, filename string) (*imagev1.ImageMetadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	metadataChan := make(chan *imagev1.ImageMetadata, 1)
	errChan := make(chan error, 1)

	go func() {
		img, err := imaging.Open(filePath)
		if err != nil {
			errChan <- fmt.Errorf("failed to open image: %w", err)
			return
		}

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			errChan <- fmt.Errorf("failed to get file info: %w", err)
			return
		}

		mimeType := getMimeType(filePath)

		metadata := &imagev1.ImageMetadata{
			Filename:    filename,
			FileSize:    fileInfo.Size(),
			MimeType:    mimeType,
			Width:       int32(img.Bounds().Dx()),
			Height:      int32(img.Bounds().Dy()),
			FilePath:    filePath,
			ImageFormat: filepath.Ext(filename)[1:],
			Tags:        generateImageTags(img),
		}

		metadataChan <- metadata
	}()

	select {
	case metadata := <-metadataChan:
		return metadata, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func GenerateThumbnail(filePath string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	thumbnailChan := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		img, err := imaging.Open(filePath)
		if err != nil {
			errChan <- fmt.Errorf("failed to open image for thumbnail: %w", err)
			return
		}

		thumbnailImg := imaging.Resize(img, 200, 0, imaging.Lanczos)

		thumbnailDir := "./uploads/thumbnails"

		thumbnailFilename := "thumb_" + filepath.Base(filePath)
		thumbnailPath := filepath.Join(thumbnailDir, thumbnailFilename)

		err = imaging.Save(thumbnailImg, thumbnailPath)
		if err != nil {
			errChan <- fmt.Errorf("failed to save thumbnail: %w", err)
			return
		}

		thumbnailChan <- thumbnailPath
	}()

	select {
	case thumbnailPath := <-thumbnailChan:
		return thumbnailPath, nil
	case err := <-errChan:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func generateImageTags(img image.Image) string {
	var tags string

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	if width > height {
		tags += " landscape"
	} else if width < height {
		tags += " portrait"
	} else {
		tags += " square"
	}

	switch {
	case width < 800 || height < 600:
		tags += " small"
	case width >= 800 && width < 1920 && height >= 600 && height < 1080:
		tags += " medium"
	case width >= 1920 && height >= 1080:
		tags += " large"
	}

	return tags
}
