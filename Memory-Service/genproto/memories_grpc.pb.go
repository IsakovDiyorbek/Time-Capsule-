// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: memories.proto

package genproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MemoryServiceClient is the client API for MemoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemoryServiceClient interface {
	AddMemory(ctx context.Context, in *AddMemoryRequest, opts ...grpc.CallOption) (*AddMemoryResponse, error)
	GetMemory(ctx context.Context, in *GetMemoryRequest, opts ...grpc.CallOption) (*GetMemoryResponse, error)
	GetAllMemories(ctx context.Context, in *GetAllMemoriesRequest, opts ...grpc.CallOption) (*GetAllMemoriesResponse, error)
	UpdateMemory(ctx context.Context, in *UpdateMemoryRequest, opts ...grpc.CallOption) (*UpdateMemoryResponse, error)
	DeleteMemory(ctx context.Context, in *DeleteMemoryRequest, opts ...grpc.CallOption) (*DeleteMemoryResponse, error)
	ShareMemory(ctx context.Context, in *ShareMemoryRequest, opts ...grpc.CallOption) (*ShareMemoryResponse, error)
}

type memoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemoryServiceClient(cc grpc.ClientConnInterface) MemoryServiceClient {
	return &memoryServiceClient{cc}
}

func (c *memoryServiceClient) AddMemory(ctx context.Context, in *AddMemoryRequest, opts ...grpc.CallOption) (*AddMemoryResponse, error) {
	out := new(AddMemoryResponse)
	err := c.cc.Invoke(ctx, "/timeline.MemoryService/AddMemory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryServiceClient) GetMemory(ctx context.Context, in *GetMemoryRequest, opts ...grpc.CallOption) (*GetMemoryResponse, error) {
	out := new(GetMemoryResponse)
	err := c.cc.Invoke(ctx, "/timeline.MemoryService/GetMemory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryServiceClient) GetAllMemories(ctx context.Context, in *GetAllMemoriesRequest, opts ...grpc.CallOption) (*GetAllMemoriesResponse, error) {
	out := new(GetAllMemoriesResponse)
	err := c.cc.Invoke(ctx, "/timeline.MemoryService/GetAllMemories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryServiceClient) UpdateMemory(ctx context.Context, in *UpdateMemoryRequest, opts ...grpc.CallOption) (*UpdateMemoryResponse, error) {
	out := new(UpdateMemoryResponse)
	err := c.cc.Invoke(ctx, "/timeline.MemoryService/UpdateMemory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryServiceClient) DeleteMemory(ctx context.Context, in *DeleteMemoryRequest, opts ...grpc.CallOption) (*DeleteMemoryResponse, error) {
	out := new(DeleteMemoryResponse)
	err := c.cc.Invoke(ctx, "/timeline.MemoryService/DeleteMemory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryServiceClient) ShareMemory(ctx context.Context, in *ShareMemoryRequest, opts ...grpc.CallOption) (*ShareMemoryResponse, error) {
	out := new(ShareMemoryResponse)
	err := c.cc.Invoke(ctx, "/timeline.MemoryService/ShareMemory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemoryServiceServer is the server API for MemoryService service.
// All implementations must embed UnimplementedMemoryServiceServer
// for forward compatibility
type MemoryServiceServer interface {
	AddMemory(context.Context, *AddMemoryRequest) (*AddMemoryResponse, error)
	GetMemory(context.Context, *GetMemoryRequest) (*GetMemoryResponse, error)
	GetAllMemories(context.Context, *GetAllMemoriesRequest) (*GetAllMemoriesResponse, error)
	UpdateMemory(context.Context, *UpdateMemoryRequest) (*UpdateMemoryResponse, error)
	DeleteMemory(context.Context, *DeleteMemoryRequest) (*DeleteMemoryResponse, error)
	ShareMemory(context.Context, *ShareMemoryRequest) (*ShareMemoryResponse, error)
	mustEmbedUnimplementedMemoryServiceServer()
}

// UnimplementedMemoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMemoryServiceServer struct {
}

func (UnimplementedMemoryServiceServer) AddMemory(context.Context, *AddMemoryRequest) (*AddMemoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMemory not implemented")
}
func (UnimplementedMemoryServiceServer) GetMemory(context.Context, *GetMemoryRequest) (*GetMemoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemory not implemented")
}
func (UnimplementedMemoryServiceServer) GetAllMemories(context.Context, *GetAllMemoriesRequest) (*GetAllMemoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMemories not implemented")
}
func (UnimplementedMemoryServiceServer) UpdateMemory(context.Context, *UpdateMemoryRequest) (*UpdateMemoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMemory not implemented")
}
func (UnimplementedMemoryServiceServer) DeleteMemory(context.Context, *DeleteMemoryRequest) (*DeleteMemoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMemory not implemented")
}
func (UnimplementedMemoryServiceServer) ShareMemory(context.Context, *ShareMemoryRequest) (*ShareMemoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareMemory not implemented")
}
func (UnimplementedMemoryServiceServer) mustEmbedUnimplementedMemoryServiceServer() {}

// UnsafeMemoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemoryServiceServer will
// result in compilation errors.
type UnsafeMemoryServiceServer interface {
	mustEmbedUnimplementedMemoryServiceServer()
}

func RegisterMemoryServiceServer(s grpc.ServiceRegistrar, srv MemoryServiceServer) {
	s.RegisterService(&MemoryService_ServiceDesc, srv)
}

func _MemoryService_AddMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).AddMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timeline.MemoryService/AddMemory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).AddMemory(ctx, req.(*AddMemoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryService_GetMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).GetMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timeline.MemoryService/GetMemory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).GetMemory(ctx, req.(*GetMemoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryService_GetAllMemories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMemoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).GetAllMemories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timeline.MemoryService/GetAllMemories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).GetAllMemories(ctx, req.(*GetAllMemoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryService_UpdateMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMemoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).UpdateMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timeline.MemoryService/UpdateMemory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).UpdateMemory(ctx, req.(*UpdateMemoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryService_DeleteMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).DeleteMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timeline.MemoryService/DeleteMemory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).DeleteMemory(ctx, req.(*DeleteMemoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryService_ShareMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareMemoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).ShareMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timeline.MemoryService/ShareMemory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).ShareMemory(ctx, req.(*ShareMemoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemoryService_ServiceDesc is the grpc.ServiceDesc for MemoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timeline.MemoryService",
	HandlerType: (*MemoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMemory",
			Handler:    _MemoryService_AddMemory_Handler,
		},
		{
			MethodName: "GetMemory",
			Handler:    _MemoryService_GetMemory_Handler,
		},
		{
			MethodName: "GetAllMemories",
			Handler:    _MemoryService_GetAllMemories_Handler,
		},
		{
			MethodName: "UpdateMemory",
			Handler:    _MemoryService_UpdateMemory_Handler,
		},
		{
			MethodName: "DeleteMemory",
			Handler:    _MemoryService_DeleteMemory_Handler,
		},
		{
			MethodName: "ShareMemory",
			Handler:    _MemoryService_ShareMemory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "memories.proto",
}
