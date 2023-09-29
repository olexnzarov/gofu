// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: pb/processmanager.proto

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
	Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*RestartReply, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
}

type processManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessManagerClient(cc grpc.ClientConnInterface) ProcessManagerClient {
	return &processManagerClient{cc}
}

func (c *processManagerClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartReply, error) {
	out := new(StartReply)
	err := c.cc.Invoke(ctx, "/processmanager.ProcessManager/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processManagerClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/processmanager.ProcessManager/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processManagerClient) Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*RestartReply, error) {
	out := new(RestartReply)
	err := c.cc.Invoke(ctx, "/processmanager.ProcessManager/Restart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processManagerClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error) {
	out := new(StopReply)
	err := c.cc.Invoke(ctx, "/processmanager.ProcessManager/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processManagerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error) {
	out := new(UpdateReply)
	err := c.cc.Invoke(ctx, "/processmanager.ProcessManager/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processManagerClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error) {
	out := new(RemoveReply)
	err := c.cc.Invoke(ctx, "/processmanager.ProcessManager/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processManagerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/processmanager.ProcessManager/Get", in, out, opts...)
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
	Restart(context.Context, *RestartRequest) (*RestartReply, error)
	Stop(context.Context, *StopRequest) (*StopReply, error)
	Update(context.Context, *UpdateRequest) (*UpdateReply, error)
	Remove(context.Context, *RemoveRequest) (*RemoveReply, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
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
func (UnimplementedProcessManagerServer) Restart(context.Context, *RestartRequest) (*RestartReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedProcessManagerServer) Stop(context.Context, *StopRequest) (*StopReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedProcessManagerServer) Update(context.Context, *UpdateRequest) (*UpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProcessManagerServer) Remove(context.Context, *RemoveRequest) (*RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedProcessManagerServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
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
		FullMethod: "/processmanager.ProcessManager/Start",
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
		FullMethod: "/processmanager.ProcessManager/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessManager_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/processmanager.ProcessManager/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServer).Restart(ctx, req.(*RestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessManager_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/processmanager.ProcessManager/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessManager_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/processmanager.ProcessManager/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessManager_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/processmanager.ProcessManager/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/processmanager.ProcessManager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProcessManager_ServiceDesc is the grpc.ServiceDesc for ProcessManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "processmanager.ProcessManager",
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
		{
			MethodName: "Restart",
			Handler:    _ProcessManager_Restart_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ProcessManager_Stop_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProcessManager_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _ProcessManager_Remove_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProcessManager_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/processmanager.proto",
}
