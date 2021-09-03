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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Register(ctx context.Context, in *presenter.RegistrationRequest, opts ...grpc.CallOption) (*presenter.RegistrationResponse, error)
	GetState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*presenter.User, error)
	PingAccountStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*presenter.UserStatusResponse, error)
	ChangePassword(ctx context.Context, in *presenter.ChangePasswordRequest, opts ...grpc.CallOption) (*presenter.StatusResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *presenter.RegistrationRequest, opts ...grpc.CallOption) (*presenter.RegistrationResponse, error) {
	out := new(presenter.RegistrationResponse)
	err := c.cc.Invoke(ctx, "/chainmetric.identity.service.UserService/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*presenter.User, error) {
	out := new(presenter.User)
	err := c.cc.Invoke(ctx, "/chainmetric.identity.service.UserService/getState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PingAccountStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*presenter.UserStatusResponse, error) {
	out := new(presenter.UserStatusResponse)
	err := c.cc.Invoke(ctx, "/chainmetric.identity.service.UserService/pingAccountStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *presenter.ChangePasswordRequest, opts ...grpc.CallOption) (*presenter.StatusResponse, error) {
	out := new(presenter.StatusResponse)
	err := c.cc.Invoke(ctx, "/chainmetric.identity.service.UserService/changePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Register(context.Context, *presenter.RegistrationRequest) (*presenter.RegistrationResponse, error)
	GetState(context.Context, *emptypb.Empty) (*presenter.User, error)
	PingAccountStatus(context.Context, *emptypb.Empty) (*presenter.UserStatusResponse, error)
	ChangePassword(context.Context, *presenter.ChangePasswordRequest) (*presenter.StatusResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *presenter.RegistrationRequest) (*presenter.RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) GetState(context.Context, *emptypb.Empty) (*presenter.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (UnimplementedUserServiceServer) PingAccountStatus(context.Context, *emptypb.Empty) (*presenter.UserStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingAccountStatus not implemented")
}
func (UnimplementedUserServiceServer) ChangePassword(context.Context, *presenter.ChangePasswordRequest) (*presenter.StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(presenter.RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmetric.identity.service.UserService/register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*presenter.RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmetric.identity.service.UserService/getState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetState(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PingAccountStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PingAccountStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmetric.identity.service.UserService/pingAccountStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PingAccountStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(presenter.ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmetric.identity.service.UserService/changePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*presenter.ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainmetric.identity.service.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "getState",
			Handler:    _UserService_GetState_Handler,
		},
		{
			MethodName: "pingAccountStatus",
			Handler:    _UserService_PingAccountStatus_Handler,
		},
		{
			MethodName: "changePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity/api/rpc/user_grpc.proto",
}
