// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/worker_server.proto

package workerServer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerClient interface {
	ReceiveTask(ctx context.Context, in *ReceiveTask, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendTaskResult(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SendTaskResult, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) ReceiveTask(ctx context.Context, in *ReceiveTask, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/workerServer.Worker/ReceiveTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) SendTaskResult(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SendTaskResult, error) {
	out := new(SendTaskResult)
	err := c.cc.Invoke(ctx, "/workerServer.Worker/SendTaskResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
// All implementations must embed UnimplementedWorkerServer
// for forward compatibility
type WorkerServer interface {
	ReceiveTask(context.Context, *ReceiveTask) (*emptypb.Empty, error)
	SendTaskResult(context.Context, *emptypb.Empty) (*SendTaskResult, error)
	mustEmbedUnimplementedWorkerServer()
}

// UnimplementedWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (UnimplementedWorkerServer) ReceiveTask(context.Context, *ReceiveTask) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveTask not implemented")
}
func (UnimplementedWorkerServer) SendTaskResult(context.Context, *emptypb.Empty) (*SendTaskResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTaskResult not implemented")
}
func (UnimplementedWorkerServer) mustEmbedUnimplementedWorkerServer() {}

// UnsafeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServer will
// result in compilation errors.
type UnsafeWorkerServer interface {
	mustEmbedUnimplementedWorkerServer()
}

func RegisterWorkerServer(s grpc.ServiceRegistrar, srv WorkerServer) {
	s.RegisterService(&Worker_ServiceDesc, srv)
}

func _Worker_ReceiveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).ReceiveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workerServer.Worker/ReceiveTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).ReceiveTask(ctx, req.(*ReceiveTask))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_SendTaskResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).SendTaskResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workerServer.Worker/SendTaskResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).SendTaskResult(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Worker_ServiceDesc is the grpc.ServiceDesc for Worker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Worker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "workerServer.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveTask",
			Handler:    _Worker_ReceiveTask_Handler,
		},
		{
			MethodName: "SendTaskResult",
			Handler:    _Worker_SendTaskResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/worker_server.proto",
}
