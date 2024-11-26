// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: image/image_service.proto

package aidosgal_image_service_v1_image_servicev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UploadImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image    []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	mi := &file_image_image_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{0}
}

func (x *UploadImageRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *UploadImageRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type UploadImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId int64 `protobuf:"varint,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	mi := &file_image_image_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{1}
}

func (x *UploadImageResponse) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

type ListImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListImagesRequest) Reset() {
	*x = ListImagesRequest{}
	mi := &file_image_image_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesRequest) ProtoMessage() {}

func (x *ListImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesRequest.ProtoReflect.Descriptor instead.
func (*ListImagesRequest) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{2}
}

type ListImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*ImageMetadata `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *ListImagesResponse) Reset() {
	*x = ListImagesResponse{}
	mi := &file_image_image_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesResponse) ProtoMessage() {}

func (x *ListImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesResponse.ProtoReflect.Descriptor instead.
func (*ListImagesResponse) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListImagesResponse) GetImages() []*ImageMetadata {
	if x != nil {
		return x.Images
	}
	return nil
}

type GetImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId int64 `protobuf:"varint,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
}

func (x *GetImageRequest) Reset() {
	*x = GetImageRequest{}
	mi := &file_image_image_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageRequest) ProtoMessage() {}

func (x *GetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageRequest.ProtoReflect.Descriptor instead.
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetImageRequest) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

type GetImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image    []byte         `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Metadata *ImageMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *GetImageResponse) Reset() {
	*x = GetImageResponse{}
	mi := &file_image_image_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageResponse) ProtoMessage() {}

func (x *GetImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageResponse.ProtoReflect.Descriptor instead.
func (*GetImageResponse) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetImageResponse) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *GetImageResponse) GetMetadata() *ImageMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type DeleteImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId int64 `protobuf:"varint,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
}

func (x *DeleteImageRequest) Reset() {
	*x = DeleteImageRequest{}
	mi := &file_image_image_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImageRequest) ProtoMessage() {}

func (x *DeleteImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImageRequest.ProtoReflect.Descriptor instead.
func (*DeleteImageRequest) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteImageRequest) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

type DeleteImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteImageResponse) Reset() {
	*x = DeleteImageResponse{}
	mi := &file_image_image_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImageResponse) ProtoMessage() {}

func (x *DeleteImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImageResponse.ProtoReflect.Descriptor instead.
func (*DeleteImageResponse) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteImageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ImageMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId       int64  `protobuf:"varint,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Filename      string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	FileSize      int64  `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	MimeType      string `protobuf:"bytes,4,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Width         int32  `protobuf:"varint,5,opt,name=width,proto3" json:"width,omitempty"`
	Height        int32  `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	CreatedAt     string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	FilePath      string `protobuf:"bytes,9,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	ThumbnailPath string `protobuf:"bytes,10,opt,name=thumbnail_path,json=thumbnailPath,proto3" json:"thumbnail_path,omitempty"`
	ImageFormat   string `protobuf:"bytes,11,opt,name=image_format,json=imageFormat,proto3" json:"image_format,omitempty"`
	Tags          string `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ImageMetadata) Reset() {
	*x = ImageMetadata{}
	mi := &file_image_image_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageMetadata) ProtoMessage() {}

func (x *ImageMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageMetadata.ProtoReflect.Descriptor instead.
func (*ImageMetadata) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{8}
}

func (x *ImageMetadata) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *ImageMetadata) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ImageMetadata) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *ImageMetadata) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *ImageMetadata) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageMetadata) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ImageMetadata) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ImageMetadata) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ImageMetadata) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *ImageMetadata) GetThumbnailPath() string {
	if x != nil {
		return x.ThumbnailPath
	}
	return ""
}

func (x *ImageMetadata) GetImageFormat() string {
	if x != nil {
		return x.ImageFormat
	}
	return ""
}

func (x *ImageMetadata) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

var File_image_image_service_proto protoreflect.FileDescriptor

var file_image_image_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x22, 0x46, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x13, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x42, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0xe7, 0x02, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x32, 0x9a, 0x02, 0x0a, 0x0c,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x18, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x19, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x61, 0x69, 0x64, 0x6f,
	0x73, 0x67, 0x61, 0x6c, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_image_image_service_proto_rawDescOnce sync.Once
	file_image_image_service_proto_rawDescData = file_image_image_service_proto_rawDesc
)

func file_image_image_service_proto_rawDescGZIP() []byte {
	file_image_image_service_proto_rawDescOnce.Do(func() {
		file_image_image_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_image_service_proto_rawDescData)
	})
	return file_image_image_service_proto_rawDescData
}

var file_image_image_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_image_image_service_proto_goTypes = []any{
	(*UploadImageRequest)(nil),  // 0: image.UploadImageRequest
	(*UploadImageResponse)(nil), // 1: image.UploadImageResponse
	(*ListImagesRequest)(nil),   // 2: image.ListImagesRequest
	(*ListImagesResponse)(nil),  // 3: image.ListImagesResponse
	(*GetImageRequest)(nil),     // 4: image.GetImageRequest
	(*GetImageResponse)(nil),    // 5: image.GetImageResponse
	(*DeleteImageRequest)(nil),  // 6: image.DeleteImageRequest
	(*DeleteImageResponse)(nil), // 7: image.DeleteImageResponse
	(*ImageMetadata)(nil),       // 8: image.ImageMetadata
}
var file_image_image_service_proto_depIdxs = []int32{
	8, // 0: image.ListImagesResponse.images:type_name -> image.ImageMetadata
	8, // 1: image.GetImageResponse.metadata:type_name -> image.ImageMetadata
	0, // 2: image.ImageService.UploadImage:input_type -> image.UploadImageRequest
	2, // 3: image.ImageService.ListImages:input_type -> image.ListImagesRequest
	4, // 4: image.ImageService.GetImage:input_type -> image.GetImageRequest
	6, // 5: image.ImageService.DeleteImage:input_type -> image.DeleteImageRequest
	1, // 6: image.ImageService.UploadImage:output_type -> image.UploadImageResponse
	3, // 7: image.ImageService.ListImages:output_type -> image.ListImagesResponse
	5, // 8: image.ImageService.GetImage:output_type -> image.GetImageResponse
	7, // 9: image.ImageService.DeleteImage:output_type -> image.DeleteImageResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_image_image_service_proto_init() }
func file_image_image_service_proto_init() {
	if File_image_image_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_image_image_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_image_service_proto_goTypes,
		DependencyIndexes: file_image_image_service_proto_depIdxs,
		MessageInfos:      file_image_image_service_proto_msgTypes,
	}.Build()
	File_image_image_service_proto = out.File
	file_image_image_service_proto_rawDesc = nil
	file_image_image_service_proto_goTypes = nil
	file_image_image_service_proto_depIdxs = nil
}
