package image

import (
	"context"

	imagev1 "github.com/aidosgal/image-processing-service/pkg/gen/go/image"
	"google.golang.org/grpc"
)

type serverAPI struct {
	imagev1.UnimplementedImageServiceServer
}

func Register(gRPC *grpc.Server) {
	imagev1.RegisterImageServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) UploadImage(ctx context.Context, req *imagev1.UploadImageRequest) (*imagev1.UploadImageResponse, error) {
	panic("implement me")
}

func (s *serverAPI) ListImages(ctx context.Context, req *imagev1.ListImagesRequest) (*imagev1.ListImagesResponse, error) {
	panic("implement me")
}

func (s *serverAPI) GetImage(ctx context.Context, req *imagev1.GetImageRequest) (*imagev1.GetImageResponse, error) {
	panic("implement me")
}

func (s *serverAPI) DeleteImage(ctx context.Context, req *imagev1.DeleteImageRequest) (*imagev1.DeleteImageResponse, error) {
	panic("implement me")
}
