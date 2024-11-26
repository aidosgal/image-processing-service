package image

import (
	"context"

	imagev1 "github.com/aidosgal/image-processing-service/pkg/gen/go/image"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImageService interface {
	UploadImage(ctx context.Context, image []byte, fileName string) (imageId int64, err error)
	ListImages(ctx context.Context) (images []*imagev1.ImageMetadata, err error)
	GetImage(ctx context.Context, image_id int64) (image []byte, metadata *imagev1.ImageMetadata, err error)
	DeleteImage(ctx context.Context, image_id int64) (is_deleted bool, err error)
}

type serverAPI struct {
	imagev1.UnimplementedImageServiceServer
	service ImageService
}

func Register(gRPC *grpc.Server, service ImageService) {
	imagev1.RegisterImageServiceServer(gRPC, &serverAPI{service: service})
}

func (s *serverAPI) UploadImage(ctx context.Context, req *imagev1.UploadImageRequest) (*imagev1.UploadImageResponse, error) {
	if req.GetImage() == nil {
		return nil, status.Error(codes.InvalidArgument, "image required")
	}

	if req.GetFilename() == "" {
		return nil, status.Error(codes.InvalidArgument, "file name required")
	}

	image_id, err := s.service.UploadImage(ctx, req.GetImage(), req.GetFilename())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &imagev1.UploadImageResponse{
		ImageId: image_id,
	}, nil
}

func (s *serverAPI) ListImages(ctx context.Context, req *imagev1.ListImagesRequest) (*imagev1.ListImagesResponse, error) {
	images, err := s.service.ListImages(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &imagev1.ListImagesResponse{
		Images: images,
	}, nil
}

func (s *serverAPI) GetImage(ctx context.Context, req *imagev1.GetImageRequest) (*imagev1.GetImageResponse, error) {
	if req.GetImageId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "image id is required")
	}

	image, metadata, err := s.service.GetImage(ctx, req.GetImageId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &imagev1.GetImageResponse{
		Image:    image,
		Metadata: metadata,
	}, nil
}

func (s *serverAPI) DeleteImage(ctx context.Context, req *imagev1.DeleteImageRequest) (*imagev1.DeleteImageResponse, error) {
	if req.GetImageId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "image id is required")
	}

	is_deleted, err := s.service.DeleteImage(ctx, req.GetImageId())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &imagev1.DeleteImageResponse{
		Success: is_deleted,
	}, nil
}
