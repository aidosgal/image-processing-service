package tests

import (
	"testing"

	imagev1 "github.com/aidosgal/image-processing-service/pkg/gen/go/image"
	suite "github.com/aidosgal/image-processing-service/tests/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListImages(t *testing.T) {
	ctx, s := suite.NewSuit(t)

	imageBytes, filename := generateTestImage()
	uploadResp, err := s.ImageServiceClient.UploadImage(ctx, &imagev1.UploadImageRequest{
		Image:    imageBytes,
		Filename: filename,
	})
	require.NoError(t, err)

	listResp, err := s.ImageServiceClient.ListImages(ctx, &imagev1.ListImagesRequest{})

	require.NoError(t, err)

	assert.NotEmpty(t, listResp.GetImages())

	found := false
	for _, img := range listResp.GetImages() {
		if img.GetImageId() == uploadResp.GetImageId() {
			found = true
			break
		}
	}
	assert.True(t, found, "Uploaded image not found in the list")
}
