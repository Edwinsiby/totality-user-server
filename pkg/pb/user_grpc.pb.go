// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: user.proto

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

const (
	User_HealthCheck_FullMethodName     = "/user.User/HealthCheck"
	User_UserDetails_FullMethodName     = "/user.User/UserDetails"
	User_UserListDetails_FullMethodName = "/user.User/UserListDetails"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	HealthCheck(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	UserDetails(ctx context.Context, in *UserDetailsRequest, opts ...grpc.CallOption) (*UserDetailsResponse, error)
	UserListDetails(ctx context.Context, in *UserListDetailsRequest, opts ...grpc.CallOption) (*UserListDetailsResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) HealthCheck(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, User_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserDetails(ctx context.Context, in *UserDetailsRequest, opts ...grpc.CallOption) (*UserDetailsResponse, error) {
	out := new(UserDetailsResponse)
	err := c.cc.Invoke(ctx, User_UserDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserListDetails(ctx context.Context, in *UserListDetailsRequest, opts ...grpc.CallOption) (*UserListDetailsResponse, error) {
	out := new(UserListDetailsResponse)
	err := c.cc.Invoke(ctx, User_UserListDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	HealthCheck(context.Context, *Request) (*Response, error)
	UserDetails(context.Context, *UserDetailsRequest) (*UserDetailsResponse, error)
	UserListDetails(context.Context, *UserListDetailsRequest) (*UserListDetailsResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) HealthCheck(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedUserServer) UserDetails(context.Context, *UserDetailsRequest) (*UserDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetails not implemented")
}
func (UnimplementedUserServer) UserListDetails(context.Context, *UserListDetailsRequest) (*UserListDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserListDetails not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).HealthCheck(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserDetails(ctx, req.(*UserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserListDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserListDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserListDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserListDetails(ctx, req.(*UserListDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _User_HealthCheck_Handler,
		},
		{
			MethodName: "UserDetails",
			Handler:    _User_UserDetails_Handler,
		},
		{
			MethodName: "UserListDetails",
			Handler:    _User_UserListDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}