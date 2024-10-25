// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: avatar.proto

package avatar_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AvatarService_SetAvatar_FullMethodName     = "/avatar.AvatarService/SetAvatar"
	AvatarService_GetAllAvatars_FullMethodName = "/avatar.AvatarService/GetAllAvatars"
	AvatarService_DeleteAvatar_FullMethodName  = "/avatar.AvatarService/DeleteAvatar"
)

// AvatarServiceClient is the client API for AvatarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AvatarServiceClient interface {
	SetAvatar(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SetAvatarIn, SetAvatarOut], error)
	GetAllAvatars(ctx context.Context, in *GetAllAvatarsIn, opts ...grpc.CallOption) (*GetAllAvatarsOut, error)
	DeleteAvatar(ctx context.Context, in *DeleteAvatarIn, opts ...grpc.CallOption) (*Avatar, error)
}

type avatarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAvatarServiceClient(cc grpc.ClientConnInterface) AvatarServiceClient {
	return &avatarServiceClient{cc}
}

func (c *avatarServiceClient) SetAvatar(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SetAvatarIn, SetAvatarOut], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AvatarService_ServiceDesc.Streams[0], AvatarService_SetAvatar_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SetAvatarIn, SetAvatarOut]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AvatarService_SetAvatarClient = grpc.ClientStreamingClient[SetAvatarIn, SetAvatarOut]

func (c *avatarServiceClient) GetAllAvatars(ctx context.Context, in *GetAllAvatarsIn, opts ...grpc.CallOption) (*GetAllAvatarsOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllAvatarsOut)
	err := c.cc.Invoke(ctx, AvatarService_GetAllAvatars_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avatarServiceClient) DeleteAvatar(ctx context.Context, in *DeleteAvatarIn, opts ...grpc.CallOption) (*Avatar, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Avatar)
	err := c.cc.Invoke(ctx, AvatarService_DeleteAvatar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AvatarServiceServer is the server API for AvatarService service.
// All implementations must embed UnimplementedAvatarServiceServer
// for forward compatibility.
type AvatarServiceServer interface {
	SetAvatar(grpc.ClientStreamingServer[SetAvatarIn, SetAvatarOut]) error
	GetAllAvatars(context.Context, *GetAllAvatarsIn) (*GetAllAvatarsOut, error)
	DeleteAvatar(context.Context, *DeleteAvatarIn) (*Avatar, error)
	mustEmbedUnimplementedAvatarServiceServer()
}

// UnimplementedAvatarServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAvatarServiceServer struct{}

func (UnimplementedAvatarServiceServer) SetAvatar(grpc.ClientStreamingServer[SetAvatarIn, SetAvatarOut]) error {
	return status.Errorf(codes.Unimplemented, "method SetAvatar not implemented")
}
func (UnimplementedAvatarServiceServer) GetAllAvatars(context.Context, *GetAllAvatarsIn) (*GetAllAvatarsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAvatars not implemented")
}
func (UnimplementedAvatarServiceServer) DeleteAvatar(context.Context, *DeleteAvatarIn) (*Avatar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAvatar not implemented")
}
func (UnimplementedAvatarServiceServer) mustEmbedUnimplementedAvatarServiceServer() {}
func (UnimplementedAvatarServiceServer) testEmbeddedByValue()                       {}

// UnsafeAvatarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AvatarServiceServer will
// result in compilation errors.
type UnsafeAvatarServiceServer interface {
	mustEmbedUnimplementedAvatarServiceServer()
}

func RegisterAvatarServiceServer(s grpc.ServiceRegistrar, srv AvatarServiceServer) {
	// If the following call pancis, it indicates UnimplementedAvatarServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AvatarService_ServiceDesc, srv)
}

func _AvatarService_SetAvatar_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AvatarServiceServer).SetAvatar(&grpc.GenericServerStream[SetAvatarIn, SetAvatarOut]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AvatarService_SetAvatarServer = grpc.ClientStreamingServer[SetAvatarIn, SetAvatarOut]

func _AvatarService_GetAllAvatars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAvatarsIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServiceServer).GetAllAvatars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvatarService_GetAllAvatars_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServiceServer).GetAllAvatars(ctx, req.(*GetAllAvatarsIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvatarService_DeleteAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAvatarIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServiceServer).DeleteAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvatarService_DeleteAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServiceServer).DeleteAvatar(ctx, req.(*DeleteAvatarIn))
	}
	return interceptor(ctx, in, info, handler)
}

// AvatarService_ServiceDesc is the grpc.ServiceDesc for AvatarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AvatarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "avatar.AvatarService",
	HandlerType: (*AvatarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllAvatars",
			Handler:    _AvatarService_GetAllAvatars_Handler,
		},
		{
			MethodName: "DeleteAvatar",
			Handler:    _AvatarService_DeleteAvatar_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SetAvatar",
			Handler:       _AvatarService_SetAvatar_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "avatar.proto",
}
