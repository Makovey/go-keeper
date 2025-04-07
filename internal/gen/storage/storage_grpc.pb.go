// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: storage.proto

package storage

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StorageService_UploadFile_FullMethodName      = "/storage.StorageService/UploadFile"
	StorageService_DownloadFile_FullMethodName    = "/storage.StorageService/DownloadFile"
	StorageService_GetUsersFile_FullMethodName    = "/storage.StorageService/GetUsersFile"
	StorageService_DeleteUsersFile_FullMethodName = "/storage.StorageService/DeleteUsersFile"
)

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadRequest, UploadResponse], error)
	DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadResponse], error)
	GetUsersFile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetUsersFileResponse, error)
	DeleteUsersFile(ctx context.Context, in *DeleteUsersFileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadRequest, UploadResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StorageService_ServiceDesc.Streams[0], StorageService_UploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadRequest, UploadResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StorageService_UploadFileClient = grpc.ClientStreamingClient[UploadRequest, UploadResponse]

func (c *storageServiceClient) DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StorageService_ServiceDesc.Streams[1], StorageService_DownloadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DownloadRequest, DownloadResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StorageService_DownloadFileClient = grpc.ServerStreamingClient[DownloadResponse]

func (c *storageServiceClient) GetUsersFile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetUsersFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUsersFileResponse)
	err := c.cc.Invoke(ctx, StorageService_GetUsersFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) DeleteUsersFile(ctx context.Context, in *DeleteUsersFileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, StorageService_DeleteUsersFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations must embed UnimplementedStorageServiceServer
// for forward compatibility.
type StorageServiceServer interface {
	UploadFile(grpc.ClientStreamingServer[UploadRequest, UploadResponse]) error
	DownloadFile(*DownloadRequest, grpc.ServerStreamingServer[DownloadResponse]) error
	GetUsersFile(context.Context, *emptypb.Empty) (*GetUsersFileResponse, error)
	DeleteUsersFile(context.Context, *DeleteUsersFileRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedStorageServiceServer()
}

// UnimplementedStorageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStorageServiceServer struct{}

func (UnimplementedStorageServiceServer) UploadFile(grpc.ClientStreamingServer[UploadRequest, UploadResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedStorageServiceServer) DownloadFile(*DownloadRequest, grpc.ServerStreamingServer[DownloadResponse]) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedStorageServiceServer) GetUsersFile(context.Context, *emptypb.Empty) (*GetUsersFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersFile not implemented")
}
func (UnimplementedStorageServiceServer) DeleteUsersFile(context.Context, *DeleteUsersFileRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUsersFile not implemented")
}
func (UnimplementedStorageServiceServer) mustEmbedUnimplementedStorageServiceServer() {}
func (UnimplementedStorageServiceServer) testEmbeddedByValue()                        {}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	// If the following call pancis, it indicates UnimplementedStorageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServiceServer).UploadFile(&grpc.GenericServerStream[UploadRequest, UploadResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StorageService_UploadFileServer = grpc.ClientStreamingServer[UploadRequest, UploadResponse]

func _StorageService_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServiceServer).DownloadFile(m, &grpc.GenericServerStream[DownloadRequest, DownloadResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StorageService_DownloadFileServer = grpc.ServerStreamingServer[DownloadResponse]

func _StorageService_GetUsersFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetUsersFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageService_GetUsersFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetUsersFile(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_DeleteUsersFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUsersFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DeleteUsersFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageService_DeleteUsersFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DeleteUsersFile(ctx, req.(*DeleteUsersFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsersFile",
			Handler:    _StorageService_GetUsersFile_Handler,
		},
		{
			MethodName: "DeleteUsersFile",
			Handler:    _StorageService_DeleteUsersFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _StorageService_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _StorageService_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "storage.proto",
}
