// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	presenter "github.com/timoth-y/chainmetric-contracts/src/identity/api/presenter"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IdentityServiceClient is the client API for IdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityServiceClient interface {
	Register(ctx context.Context, in *presenter.RegistrationRequest, opts ...grpc.CallOption) (*presenter.RegistrationResponse, error)
	Enroll(ctx context.Context, in *presenter.EnrollmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type identityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityServiceClient(cc grpc.ClientConnInterface) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) Register(ctx context.Context, in *presenter.RegistrationRequest, opts ...grpc.CallOption) (*presenter.RegistrationResponse, error) {
	out := new(presenter.RegistrationResponse)
	err := c.cc.Invoke(ctx, "/chainmetric.identity.service.IdentityService/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) Enroll(ctx context.Context, in *presenter.EnrollmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chainmetric.identity.service.IdentityService/enroll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServiceServer is the server API for IdentityService service.
// All implementations must embed UnimplementedIdentityServiceServer
// for forward compatibility
type IdentityServiceServer interface {
	Register(context.Context, *presenter.RegistrationRequest) (*presenter.RegistrationResponse, error)
	Enroll(context.Context, *presenter.EnrollmentRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedIdentityServiceServer()
}

// UnimplementedIdentityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityServiceServer struct {
}

func (UnimplementedIdentityServiceServer) Register(context.Context, *presenter.RegistrationRequest) (*presenter.RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedIdentityServiceServer) Enroll(context.Context, *presenter.EnrollmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enroll not implemented")
}
func (UnimplementedIdentityServiceServer) mustEmbedUnimplementedIdentityServiceServer() {}

// UnsafeIdentityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServiceServer will
// result in compilation errors.
type UnsafeIdentityServiceServer interface {
	mustEmbedUnimplementedIdentityServiceServer()
}

func RegisterIdentityServiceServer(s grpc.ServiceRegistrar, srv IdentityServiceServer) {
	s.RegisterService(&IdentityService_ServiceDesc, srv)
}

func _IdentityService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(presenter.RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmetric.identity.service.IdentityService/register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).Register(ctx, req.(*presenter.RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_Enroll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(presenter.EnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).Enroll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmetric.identity.service.IdentityService/enroll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).Enroll(ctx, req.(*presenter.EnrollmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityService_ServiceDesc is the grpc.ServiceDesc for IdentityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainmetric.identity.service.IdentityService",
	HandlerType: (*IdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _IdentityService_Register_Handler,
		},
		{
			MethodName: "enroll",
			Handler:    _IdentityService_Enroll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity/api/rpc/identity.proto",
}
