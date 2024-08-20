// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: timeline.proto

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

// TimeLineServiceClient is the client API for TimeLineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeLineServiceClient interface {
	AddTimeLine(ctx context.Context, in *AddTimeLineRequest, opts ...grpc.CallOption) (*AddTimeLineResponse, error)
	GetEvent(ctx context.Context, in *GetUserEventsRequest, opts ...grpc.CallOption) (*GetUserEventsResponse, error)
	SearchTimeLine(ctx context.Context, in *SearchTimeLineRequest, opts ...grpc.CallOption) (*SearchTimeLineResponse, error)
	DeleteTimeLine(ctx context.Context, in *DeleteTimeLineRequest, opts ...grpc.CallOption) (*DeleteTimeLineResponse, error)
}

type timeLineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeLineServiceClient(cc grpc.ClientConnInterface) TimeLineServiceClient {
	return &timeLineServiceClient{cc}
}

func (c *timeLineServiceClient) AddTimeLine(ctx context.Context, in *AddTimeLineRequest, opts ...grpc.CallOption) (*AddTimeLineResponse, error) {
	out := new(AddTimeLineResponse)
	err := c.cc.Invoke(ctx, "/github.com/Time-Capsule/Auth-Service
.TimeLineService/AddTimeLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeLineServiceClient) GetEvent(ctx context.Context, in *GetUserEventsRequest, opts ...grpc.CallOption) (*GetUserEventsResponse, error) {
	out := new(GetUserEventsResponse)
	err := c.cc.Invoke(ctx, "/github.com/Time-Capsule/Auth-Service
.TimeLineService/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeLineServiceClient) SearchTimeLine(ctx context.Context, in *SearchTimeLineRequest, opts ...grpc.CallOption) (*SearchTimeLineResponse, error) {
	out := new(SearchTimeLineResponse)
	err := c.cc.Invoke(ctx, "/github.com/Time-Capsule/Auth-Service
.TimeLineService/SearchTimeLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeLineServiceClient) DeleteTimeLine(ctx context.Context, in *DeleteTimeLineRequest, opts ...grpc.CallOption) (*DeleteTimeLineResponse, error) {
	out := new(DeleteTimeLineResponse)
	err := c.cc.Invoke(ctx, "/github.com/Time-Capsule/Auth-Service
.TimeLineService/DeleteTimeLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeLineServiceServer is the server API for TimeLineService service.
// All implementations must embed UnimplementedTimeLineServiceServer
// for forward compatibility
type TimeLineServiceServer interface {
	AddTimeLine(context.Context, *AddTimeLineRequest) (*AddTimeLineResponse, error)
	GetEvent(context.Context, *GetUserEventsRequest) (*GetUserEventsResponse, error)
	SearchTimeLine(context.Context, *SearchTimeLineRequest) (*SearchTimeLineResponse, error)
	DeleteTimeLine(context.Context, *DeleteTimeLineRequest) (*DeleteTimeLineResponse, error)
	mustEmbedUnimplementedTimeLineServiceServer()
}

// UnimplementedTimeLineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimeLineServiceServer struct {
}

func (UnimplementedTimeLineServiceServer) AddTimeLine(context.Context, *AddTimeLineRequest) (*AddTimeLineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTimeLine not implemented")
}
func (UnimplementedTimeLineServiceServer) GetEvent(context.Context, *GetUserEventsRequest) (*GetUserEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedTimeLineServiceServer) SearchTimeLine(context.Context, *SearchTimeLineRequest) (*SearchTimeLineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTimeLine not implemented")
}
func (UnimplementedTimeLineServiceServer) DeleteTimeLine(context.Context, *DeleteTimeLineRequest) (*DeleteTimeLineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTimeLine not implemented")
}
func (UnimplementedTimeLineServiceServer) mustEmbedUnimplementedTimeLineServiceServer() {}

// UnsafeTimeLineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeLineServiceServer will
// result in compilation errors.
type UnsafeTimeLineServiceServer interface {
	mustEmbedUnimplementedTimeLineServiceServer()
}

func RegisterTimeLineServiceServer(s grpc.ServiceRegistrar, srv TimeLineServiceServer) {
	s.RegisterService(&TimeLineService_ServiceDesc, srv)
}

func _TimeLineService_AddTimeLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTimeLineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeLineServiceServer).AddTimeLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com/Time-Capsule/Auth-Service
.TimeLineService/AddTimeLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeLineServiceServer).AddTimeLine(ctx, req.(*AddTimeLineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeLineService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeLineServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com/Time-Capsule/Auth-Service
.TimeLineService/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeLineServiceServer).GetEvent(ctx, req.(*GetUserEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeLineService_SearchTimeLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTimeLineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeLineServiceServer).SearchTimeLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com/Time-Capsule/Auth-Service
.TimeLineService/SearchTimeLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeLineServiceServer).SearchTimeLine(ctx, req.(*SearchTimeLineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeLineService_DeleteTimeLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTimeLineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeLineServiceServer).DeleteTimeLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com/Time-Capsule/Auth-Service
.TimeLineService/DeleteTimeLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeLineServiceServer).DeleteTimeLine(ctx, req.(*DeleteTimeLineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeLineService_ServiceDesc is the grpc.ServiceDesc for TimeLineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeLineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com/Time-Capsule/Auth-Service
.TimeLineService",
	HandlerType: (*TimeLineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTimeLine",
			Handler:    _TimeLineService_AddTimeLine_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _TimeLineService_GetEvent_Handler,
		},
		{
			MethodName: "SearchTimeLine",
			Handler:    _TimeLineService_SearchTimeLine_Handler,
		},
		{
			MethodName: "DeleteTimeLine",
			Handler:    _TimeLineService_DeleteTimeLine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timeline.proto",
}
