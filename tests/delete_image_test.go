package tests

import (
	"testing"

	imagev1 "github.com/aidosgal/image-processing-service/pkg/gen/go/image"
	suite "github.com/aidosgal/image-processing-service/tests/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteImage(t *testing.T) {
	ctx, s := suite.NewSuit(t)

	imageBytes, filename := generateTestImage()
	uploadResp, err := s.ImageServiceClient.UploadImage(ctx, &imagev1.UploadImageRequest{
		Image:    imageBytes,
		Filename: filename,
	})
	require.NoError(t, err)

	deleteResp, err := s.ImageServiceClient.DeleteImage(ctx, &imagev1.DeleteImageRequest{
		ImageId: uploadResp.GetImageId(),
	})

	require.NoError(t, err)

	assert.True(t, deleteResp.GetSuccess())

	_, err = s.ImageServiceClient.GetImage(ctx, &imagev1.GetImageRequest{
		ImageId: uploadResp.GetImageId(),
	})

	assert.Error(t, err)
}
