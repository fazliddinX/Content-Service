// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.1
// source: auth_service.proto

package auth_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AuthService_Register_FullMethodName         = "/protos.AuthService/Register"
	AuthService_Login_FullMethodName            = "/protos.AuthService/Login"
	AuthService_GetProfile_FullMethodName       = "/protos.AuthService/GetProfile"
	AuthService_UpdateProfile_FullMethodName    = "/protos.AuthService/UpdateProfile"
	AuthService_GetUsers_FullMethodName         = "/protos.AuthService/GetUsers"
	AuthService_Delete_FullMethodName           = "/protos.AuthService/Delete"
	AuthService_PasswordRecovery_FullMethodName = "/protos.AuthService/PasswordRecovery"
	AuthService_TokenRenewal_FullMethodName     = "/protos.AuthService/TokenRenewal"
	AuthService_Logout_FullMethodName           = "/protos.AuthService/Logout"
	AuthService_ActivityProfile_FullMethodName  = "/protos.AuthService/ActivityProfile"
	AuthService_Follow_FullMethodName           = "/protos.AuthService/Follow"
	AuthService_Unfollow_FullMethodName         = "/protos.AuthService/Unfollow"
	AuthService_GetFollowers_FullMethodName     = "/protos.AuthService/GetFollowers"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Register(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*RegisterUserRes, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Tokens, error)
	GetProfile(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Profile, error)
	UpdateProfile(ctx context.Context, in *UpdateUser, opts ...grpc.CallOption) (*Profile, error)
	GetUsers(ctx context.Context, in *FilterGet, opts ...grpc.CallOption) (*Users, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Success, error)
	PasswordRecovery(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Success, error)
	TokenRenewal(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*Tokens, error)
	Logout(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Success, error)
	ActivityProfile(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserActivities, error)
	Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error)
	Unfollow(ctx context.Context, in *FollowingId, opts ...grpc.CallOption) (*Success, error)
	GetFollowers(ctx context.Context, in *FilterFollowers, opts ...grpc.CallOption) (*Followers, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*RegisterUserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterUserRes)
	err := c.cc.Invoke(ctx, AuthService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Tokens, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Tokens)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetProfile(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Profile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Profile)
	err := c.cc.Invoke(ctx, AuthService_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateProfile(ctx context.Context, in *UpdateUser, opts ...grpc.CallOption) (*Profile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Profile)
	err := c.cc.Invoke(ctx, AuthService_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUsers(ctx context.Context, in *FilterGet, opts ...grpc.CallOption) (*Users, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Users)
	err := c.cc.Invoke(ctx, AuthService_GetUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Success, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Success)
	err := c.cc.Invoke(ctx, AuthService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) PasswordRecovery(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Success, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Success)
	err := c.cc.Invoke(ctx, AuthService_PasswordRecovery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) TokenRenewal(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*Tokens, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Tokens)
	err := c.cc.Invoke(ctx, AuthService_TokenRenewal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Success, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Success)
	err := c.cc.Invoke(ctx, AuthService_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ActivityProfile(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserActivities, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserActivities)
	err := c.cc.Invoke(ctx, AuthService_ActivityProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FollowResponse)
	err := c.cc.Invoke(ctx, AuthService_Follow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Unfollow(ctx context.Context, in *FollowingId, opts ...grpc.CallOption) (*Success, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Success)
	err := c.cc.Invoke(ctx, AuthService_Unfollow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetFollowers(ctx context.Context, in *FilterFollowers, opts ...grpc.CallOption) (*Followers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Followers)
	err := c.cc.Invoke(ctx, AuthService_GetFollowers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Register(context.Context, *RegisterUserReq) (*RegisterUserRes, error)
	Login(context.Context, *LoginRequest) (*Tokens, error)
	GetProfile(context.Context, *Id) (*Profile, error)
	UpdateProfile(context.Context, *UpdateUser) (*Profile, error)
	GetUsers(context.Context, *FilterGet) (*Users, error)
	Delete(context.Context, *Id) (*Success, error)
	PasswordRecovery(context.Context, *Email) (*Success, error)
	TokenRenewal(context.Context, *RefreshToken) (*Tokens, error)
	Logout(context.Context, *Id) (*Success, error)
	ActivityProfile(context.Context, *Id) (*UserActivities, error)
	Follow(context.Context, *FollowRequest) (*FollowResponse, error)
	Unfollow(context.Context, *FollowingId) (*Success, error)
	GetFollowers(context.Context, *FilterFollowers) (*Followers, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Register(context.Context, *RegisterUserReq) (*RegisterUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*Tokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) GetProfile(context.Context, *Id) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedAuthServiceServer) UpdateProfile(context.Context, *UpdateUser) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedAuthServiceServer) GetUsers(context.Context, *FilterGet) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedAuthServiceServer) Delete(context.Context, *Id) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAuthServiceServer) PasswordRecovery(context.Context, *Email) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PasswordRecovery not implemented")
}
func (UnimplementedAuthServiceServer) TokenRenewal(context.Context, *RefreshToken) (*Tokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenRenewal not implemented")
}
func (UnimplementedAuthServiceServer) Logout(context.Context, *Id) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthServiceServer) ActivityProfile(context.Context, *Id) (*UserActivities, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityProfile not implemented")
}
func (UnimplementedAuthServiceServer) Follow(context.Context, *FollowRequest) (*FollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedAuthServiceServer) Unfollow(context.Context, *FollowingId) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unfollow not implemented")
}
func (UnimplementedAuthServiceServer) GetFollowers(context.Context, *FilterFollowers) (*Followers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowers not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetProfile(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateProfile(ctx, req.(*UpdateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUsers(ctx, req.(*FilterGet))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_PasswordRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PasswordRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_PasswordRecovery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PasswordRecovery(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_TokenRenewal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).TokenRenewal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_TokenRenewal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).TokenRenewal(ctx, req.(*RefreshToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ActivityProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ActivityProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ActivityProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ActivityProfile(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Follow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Follow(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Unfollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowingId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Unfollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Unfollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Unfollow(ctx, req.(*FollowingId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterFollowers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetFollowers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetFollowers(ctx, req.(*FilterFollowers))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _AuthService_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _AuthService_UpdateProfile_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _AuthService_GetUsers_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AuthService_Delete_Handler,
		},
		{
			MethodName: "PasswordRecovery",
			Handler:    _AuthService_PasswordRecovery_Handler,
		},
		{
			MethodName: "TokenRenewal",
			Handler:    _AuthService_TokenRenewal_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
		{
			MethodName: "ActivityProfile",
			Handler:    _AuthService_ActivityProfile_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _AuthService_Follow_Handler,
		},
		{
			MethodName: "Unfollow",
			Handler:    _AuthService_Unfollow_Handler,
		},
		{
			MethodName: "GetFollowers",
			Handler:    _AuthService_GetFollowers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}