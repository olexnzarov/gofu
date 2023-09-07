// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: pb/process_manager.proto

package pb

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

// ProcessManagerClient is the client API for ProcessManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessManagerClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
}

type processManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessManagerClient(cc grpc.ClientConnInterface) ProcessManagerClient {
	return &processManagerClient{cc}
}

func (c *processManagerClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartReply, error) {
	out := new(StartReply)
	err := c.cc.Invoke(ctx, "/process_manager.ProcessManager/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processManagerClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/process_manager.ProcessManager/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessManagerServer is the server API for ProcessManager service.
// All implementations must embed UnimplementedProcessManagerServer
// for forward compatibility
type ProcessManagerServer interface {
	Start(context.Context, *StartRequest) (*StartReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	mustEmbedUnimplementedProcessManagerServer()
}

// UnimplementedProcessManagerServer must be embedded to have forward compatible implementations.
type UnimplementedProcessManagerServer struct {
}

func (UnimplementedProcessManagerServer) Start(context.Context, *StartRequest) (*StartReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedProcessManagerServer) List(context.Context, *ListRequest) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProcessManagerServer) mustEmbedUnimplementedProcessManagerServer() {}

// UnsafeProcessManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessManagerServer will
// result in compilation errors.
type UnsafeProcessManagerServer interface {
	mustEmbedUnimplementedProcessManagerServer()
}

func RegisterProcessManagerServer(s grpc.ServiceRegistrar, srv ProcessManagerServer) {
	s.RegisterService(&ProcessManager_ServiceDesc, srv)
}

func _ProcessManager_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/process_manager.ProcessManager/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessManager_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/process_manager.ProcessManager/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProcessManager_ServiceDesc is the grpc.ServiceDesc for ProcessManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "process_manager.ProcessManager",
	HandlerType: (*ProcessManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _ProcessManager_Start_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ProcessManager_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/process_manager.proto",
}