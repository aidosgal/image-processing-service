package tests

import (
	"bytes"
	"crypto/rand"
	"image"
	"image/jpeg"
	"testing"

	imagev1 "github.com/aidosgal/image-processing-service/pkg/gen/go/image"
	suite "github.com/aidosgal/image-processing-service/tests/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func generateTestImage() ([]byte, string) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	var buf bytes.Buffer

	err := jpeg.Encode(&buf, img, nil)
	if err != nil {
		panic(err)
	}

	return buf.Bytes(), "test_image.jpg"
}

func generateRandomImage(size int) ([]byte, string) {
	randomBytes := make([]byte, size)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	return randomBytes, "random_image.bin"
}

func TestUploadImage_HappyPath(t *testing.T) {
	ctx, s := suite.NewSuit(t)

	imageBytes, filename := generateTestImage()

	uploadResp, err := s.ImageServiceClient.UploadImage(ctx, &imagev1.UploadImageRequest{
		Image:    imageBytes,
		Filename: filename,
	})

	require.NoError(t, err)

	assert.Greater(t, uploadResp.GetImageId(), int64(0))
}

func TestUploadImage_LargeImage(t *testing.T) {
	ctx, s := suite.NewSuit(t)

	largeImageBytes := make([]byte, 5*1024*1024)
	_, err := rand.Read(largeImageBytes)
	require.NoError(t, err)

	uploadResp, err := s.ImageServiceClient.UploadImage(ctx, &imagev1.UploadImageRequest{
		Image:    largeImageBytes,
		Filename: "large_image.bin",
	})

	require.Error(t, err)
	assert.Equal(t, uploadResp.GetImageId(), int64(0))
	assert.Error(t, err)
}

func TestUploadImage_InvalidImage(t *testing.T) {
	testCases := []struct {
		name     string
		image    []byte
		filename string
	}{
		{
			name:     "Empty Image",
			image:    []byte{},
			filename: "empty.jpg",
		},
		{
			name:     "Random Bytes",
			image:    []byte{1, 2, 3, 4, 5},
			filename: "random.bin",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, s := suite.NewSuit(t)

			_, err := s.ImageServiceClient.UploadImage(ctx, &imagev1.UploadImageRequest{
				Image:    tc.image,
				Filename: tc.filename,
			})

			assert.Error(t, err)
		})
	}
}

func TestUploadMultipleImages(t *testing.T) {
	ctx, s := suite.NewSuit(t)

	uploadResults := make(chan struct {
		ImageID int64
		Err     error
	}, 5)

	for i := 0; i < 5; i++ {
		go func() {
			imageBytes, filename := generateTestImage()
			uploadResp, err := s.ImageServiceClient.UploadImage(ctx, &imagev1.UploadImageRequest{
				Image:    imageBytes,
				Filename: filename,
			})

			uploadResults <- struct {
				ImageID int64
				Err     error
			}{
				ImageID: uploadResp.GetImageId(),
				Err:     err,
			}
		}()
	}

	for i := 0; i < 5; i++ {
		result := <-uploadResults
		require.NoError(t, result.Err)
		assert.Greater(t, result.ImageID, int64(0))
	}
}
