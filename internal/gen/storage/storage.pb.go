// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: storage.proto

package storage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TextType int32

const (
	TextType_unsecure TextType = 0
	TextType_secure   TextType = 1
)

// Enum value maps for TextType.
var (
	TextType_name = map[int32]string{
		0: "unsecure",
		1: "secure",
	}
	TextType_value = map[string]int32{
		"unsecure": 0,
		"secure":   1,
	}
)

func (x TextType) Enum() *TextType {
	p := new(TextType)
	*p = x
	return p
}

func (x TextType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TextType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_proto_enumTypes[0].Descriptor()
}

func (TextType) Type() protoreflect.EnumType {
	return &file_storage_proto_enumTypes[0]
}

func (x TextType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TextType.Descriptor instead.
func (TextType) EnumDescriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{0}
}

type UploadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChunkData     []byte                 `protobuf:"bytes,1,opt,name=chunk_data,json=chunkData,proto3" json:"chunk_data,omitempty"`
	FileName      string                 `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	mi := &file_storage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

func (x *UploadRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type UploadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	mi := &file_storage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DownloadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	mi := &file_storage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{2}
}

func (x *DownloadRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DownloadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChunkData     []byte                 `protobuf:"bytes,1,opt,name=chunk_data,json=chunkData,proto3" json:"chunk_data,omitempty"`
	FileName      string                 `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	mi := &file_storage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResponse.ProtoReflect.Descriptor instead.
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadResponse) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

func (x *DownloadResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type GetUsersFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Files         []*UsersFile           `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUsersFileResponse) Reset() {
	*x = GetUsersFileResponse{}
	mi := &file_storage_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUsersFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersFileResponse) ProtoMessage() {}

func (x *GetUsersFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersFileResponse.ProtoReflect.Descriptor instead.
func (*GetUsersFileResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{4}
}

func (x *GetUsersFileResponse) GetFiles() []*UsersFile {
	if x != nil {
		return x.Files
	}
	return nil
}

type UsersFile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	FileName      string                 `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileSize      string                 `protobuf:"bytes,3,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsersFile) Reset() {
	*x = UsersFile{}
	mi := &file_storage_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersFile) ProtoMessage() {}

func (x *UsersFile) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersFile.ProtoReflect.Descriptor instead.
func (*UsersFile) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{5}
}

func (x *UsersFile) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UsersFile) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UsersFile) GetFileSize() string {
	if x != nil {
		return x.FileSize
	}
	return ""
}

func (x *UsersFile) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type DeleteUsersFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileId        string                 `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUsersFileRequest) Reset() {
	*x = DeleteUsersFileRequest{}
	mi := &file_storage_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUsersFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUsersFileRequest) ProtoMessage() {}

func (x *DeleteUsersFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUsersFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteUsersFileRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteUsersFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DeleteUsersFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type UploadPlainTextTypeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       string                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Type          TextType               `protobuf:"varint,2,opt,name=type,proto3,enum=storage.TextType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadPlainTextTypeRequest) Reset() {
	*x = UploadPlainTextTypeRequest{}
	mi := &file_storage_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadPlainTextTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadPlainTextTypeRequest) ProtoMessage() {}

func (x *UploadPlainTextTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadPlainTextTypeRequest.ProtoReflect.Descriptor instead.
func (*UploadPlainTextTypeRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{7}
}

func (x *UploadPlainTextTypeRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UploadPlainTextTypeRequest) GetType() TextType {
	if x != nil {
		return x.Type
	}
	return TextType_unsecure
}

type UploadPlainTextTypeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadPlainTextTypeResponse) Reset() {
	*x = UploadPlainTextTypeResponse{}
	mi := &file_storage_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadPlainTextTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadPlainTextTypeResponse) ProtoMessage() {}

func (x *UploadPlainTextTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadPlainTextTypeResponse.ProtoReflect.Descriptor instead.
func (*UploadPlainTextTypeResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{8}
}

func (x *UploadPlainTextTypeResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

var File_storage_proto protoreflect.FileDescriptor

var file_storage_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x2a,
	0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x10, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x99, 0x01, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4e, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x1a, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a, 0x1b, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x2a, 0x24, 0x0a, 0x08, 0x54, 0x65, 0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0c, 0x0a, 0x08, 0x75, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x10, 0x01, 0x32, 0x8d, 0x03, 0x0a, 0x0e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x45, 0x0a,
	0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1f,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x60, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50,
	0x6c, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x6b, 0x6f, 0x76, 0x65, 0x79, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_storage_proto_rawDescOnce sync.Once
	file_storage_proto_rawDescData []byte
)

func file_storage_proto_rawDescGZIP() []byte {
	file_storage_proto_rawDescOnce.Do(func() {
		file_storage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_proto_rawDesc), len(file_storage_proto_rawDesc)))
	})
	return file_storage_proto_rawDescData
}

var file_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_storage_proto_goTypes = []any{
	(TextType)(0),                       // 0: storage.TextType
	(*UploadRequest)(nil),               // 1: storage.UploadRequest
	(*UploadResponse)(nil),              // 2: storage.UploadResponse
	(*DownloadRequest)(nil),             // 3: storage.DownloadRequest
	(*DownloadResponse)(nil),            // 4: storage.DownloadResponse
	(*GetUsersFileResponse)(nil),        // 5: storage.GetUsersFileResponse
	(*UsersFile)(nil),                   // 6: storage.UsersFile
	(*DeleteUsersFileRequest)(nil),      // 7: storage.DeleteUsersFileRequest
	(*UploadPlainTextTypeRequest)(nil),  // 8: storage.UploadPlainTextTypeRequest
	(*UploadPlainTextTypeResponse)(nil), // 9: storage.UploadPlainTextTypeResponse
	(*timestamppb.Timestamp)(nil),       // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),               // 11: google.protobuf.Empty
}
var file_storage_proto_depIdxs = []int32{
	6,  // 0: storage.GetUsersFileResponse.files:type_name -> storage.UsersFile
	10, // 1: storage.UsersFile.created_at:type_name -> google.protobuf.Timestamp
	0,  // 2: storage.UploadPlainTextTypeRequest.type:type_name -> storage.TextType
	1,  // 3: storage.StorageService.UploadFile:input_type -> storage.UploadRequest
	3,  // 4: storage.StorageService.DownloadFile:input_type -> storage.DownloadRequest
	11, // 5: storage.StorageService.GetUsersFile:input_type -> google.protobuf.Empty
	7,  // 6: storage.StorageService.DeleteUsersFile:input_type -> storage.DeleteUsersFileRequest
	8,  // 7: storage.StorageService.UploadPlainTextType:input_type -> storage.UploadPlainTextTypeRequest
	2,  // 8: storage.StorageService.UploadFile:output_type -> storage.UploadResponse
	4,  // 9: storage.StorageService.DownloadFile:output_type -> storage.DownloadResponse
	5,  // 10: storage.StorageService.GetUsersFile:output_type -> storage.GetUsersFileResponse
	11, // 11: storage.StorageService.DeleteUsersFile:output_type -> google.protobuf.Empty
	9,  // 12: storage.StorageService.UploadPlainTextType:output_type -> storage.UploadPlainTextTypeResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_storage_proto_init() }
func file_storage_proto_init() {
	if File_storage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_proto_rawDesc), len(file_storage_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_proto_goTypes,
		DependencyIndexes: file_storage_proto_depIdxs,
		EnumInfos:         file_storage_proto_enumTypes,
		MessageInfos:      file_storage_proto_msgTypes,
	}.Build()
	File_storage_proto = out.File
	file_storage_proto_goTypes = nil
	file_storage_proto_depIdxs = nil
}
