// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: experiment.proto

package tutorial

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

// ExperimentClient is the client API for Experiment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExperimentClient interface {
	// Sends a greeting
	Test1(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Test1Res, error)
	Test2(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Test23Res, error)
	Test3(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Test23Res, error)
}

type experimentClient struct {
	cc grpc.ClientConnInterface
}

func NewExperimentClient(cc grpc.ClientConnInterface) ExperimentClient {
	return &experimentClient{cc}
}

func (c *experimentClient) Test1(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Test1Res, error) {
	out := new(Test1Res)
	err := c.cc.Invoke(ctx, "/Experiment/Test1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentClient) Test2(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Test23Res, error) {
	out := new(Test23Res)
	err := c.cc.Invoke(ctx, "/Experiment/Test2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentClient) Test3(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Test23Res, error) {
	out := new(Test23Res)
	err := c.cc.Invoke(ctx, "/Experiment/Test3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExperimentServer is the server API for Experiment service.
// All implementations must embed UnimplementedExperimentServer
// for forward compatibility
type ExperimentServer interface {
	// Sends a greeting
	Test1(context.Context, *Empty) (*Test1Res, error)
	Test2(context.Context, *Empty) (*Test23Res, error)
	Test3(context.Context, *Empty) (*Test23Res, error)
	mustEmbedUnimplementedExperimentServer()
}

// UnimplementedExperimentServer must be embedded to have forward compatible implementations.
type UnimplementedExperimentServer struct {
}

func (UnimplementedExperimentServer) Test1(context.Context, *Empty) (*Test1Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test1 not implemented")
}
func (UnimplementedExperimentServer) Test2(context.Context, *Empty) (*Test23Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test2 not implemented")
}
func (UnimplementedExperimentServer) Test3(context.Context, *Empty) (*Test23Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test3 not implemented")
}
func (UnimplementedExperimentServer) mustEmbedUnimplementedExperimentServer() {}

// UnsafeExperimentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExperimentServer will
// result in compilation errors.
type UnsafeExperimentServer interface {
	mustEmbedUnimplementedExperimentServer()
}

func RegisterExperimentServer(s grpc.ServiceRegistrar, srv ExperimentServer) {
	s.RegisterService(&Experiment_ServiceDesc, srv)
}

func _Experiment_Test1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServer).Test1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Experiment/Test1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServer).Test1(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Experiment_Test2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServer).Test2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Experiment/Test2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServer).Test2(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Experiment_Test3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServer).Test3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Experiment/Test3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServer).Test3(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Experiment_ServiceDesc is the grpc.ServiceDesc for Experiment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Experiment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Experiment",
	HandlerType: (*ExperimentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test1",
			Handler:    _Experiment_Test1_Handler,
		},
		{
			MethodName: "Test2",
			Handler:    _Experiment_Test2_Handler,
		},
		{
			MethodName: "Test3",
			Handler:    _Experiment_Test3_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "experiment.proto",
}