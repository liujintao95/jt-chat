// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Login(ctx context.Context, in *LoginIn, opts ...grpc.CallOption) (*LoginOut, error)
	Register(ctx context.Context, in *RegisterIn, opts ...grpc.CallOption) (*RegisterOut, error)
	GetUserList(ctx context.Context, in *GetListIn, opts ...grpc.CallOption) (*GetListOut, error)
	UpdateUser(ctx context.Context, in *UpdateIn, opts ...grpc.CallOption) (*UpdateOut, error)
	CreateGroup(ctx context.Context, in *CreateGroupIn, opts ...grpc.CallOption) (*CreateGroupOut, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupIn, opts ...grpc.CallOption) (*UpdateGroupOut, error)
	DeleteUserGroupMap(ctx context.Context, in *DeleteUserGroupMapIn, opts ...grpc.CallOption) (*DeleteUserGroupMapOut, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupIn, opts ...grpc.CallOption) (*DeleteGroupOut, error)
	GetContactList(ctx context.Context, in *GetContactListIn, opts ...grpc.CallOption) (*GetContactListOut, error)
	DeleteContact(ctx context.Context, in *DeleteContactIn, opts ...grpc.CallOption) (*DeleteContactOut, error)
	GetContactApplicationList(ctx context.Context, in *GetContactApplicationListIn, opts ...grpc.CallOption) (*GetContactApplicationListOut, error)
	GetGroupContactApplicationList(ctx context.Context, in *GetGroupContactApplicationListIn, opts ...grpc.CallOption) (*GetGroupContactApplicationListOut, error)
	CreateContactApplication(ctx context.Context, in *CreateContactApplicationIn, opts ...grpc.CallOption) (*CreateContactApplicationOut, error)
	UpdateContactApplication(ctx context.Context, in *UpdateContactApplicationIn, opts ...grpc.CallOption) (*UpdateContactApplicationOut, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginIn, opts ...grpc.CallOption) (*LoginOut, error) {
	out := new(LoginOut)
	err := c.cc.Invoke(ctx, "/pb.user/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Register(ctx context.Context, in *RegisterIn, opts ...grpc.CallOption) (*RegisterOut, error) {
	out := new(RegisterOut)
	err := c.cc.Invoke(ctx, "/pb.user/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserList(ctx context.Context, in *GetListIn, opts ...grpc.CallOption) (*GetListOut, error) {
	out := new(GetListOut)
	err := c.cc.Invoke(ctx, "/pb.user/getUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateIn, opts ...grpc.CallOption) (*UpdateOut, error) {
	out := new(UpdateOut)
	err := c.cc.Invoke(ctx, "/pb.user/updateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateGroup(ctx context.Context, in *CreateGroupIn, opts ...grpc.CallOption) (*CreateGroupOut, error) {
	out := new(CreateGroupOut)
	err := c.cc.Invoke(ctx, "/pb.user/createGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateGroup(ctx context.Context, in *UpdateGroupIn, opts ...grpc.CallOption) (*UpdateGroupOut, error) {
	out := new(UpdateGroupOut)
	err := c.cc.Invoke(ctx, "/pb.user/updateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUserGroupMap(ctx context.Context, in *DeleteUserGroupMapIn, opts ...grpc.CallOption) (*DeleteUserGroupMapOut, error) {
	out := new(DeleteUserGroupMapOut)
	err := c.cc.Invoke(ctx, "/pb.user/deleteUserGroupMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteGroup(ctx context.Context, in *DeleteGroupIn, opts ...grpc.CallOption) (*DeleteGroupOut, error) {
	out := new(DeleteGroupOut)
	err := c.cc.Invoke(ctx, "/pb.user/deleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetContactList(ctx context.Context, in *GetContactListIn, opts ...grpc.CallOption) (*GetContactListOut, error) {
	out := new(GetContactListOut)
	err := c.cc.Invoke(ctx, "/pb.user/getContactList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteContact(ctx context.Context, in *DeleteContactIn, opts ...grpc.CallOption) (*DeleteContactOut, error) {
	out := new(DeleteContactOut)
	err := c.cc.Invoke(ctx, "/pb.user/deleteContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetContactApplicationList(ctx context.Context, in *GetContactApplicationListIn, opts ...grpc.CallOption) (*GetContactApplicationListOut, error) {
	out := new(GetContactApplicationListOut)
	err := c.cc.Invoke(ctx, "/pb.user/getContactApplicationList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetGroupContactApplicationList(ctx context.Context, in *GetGroupContactApplicationListIn, opts ...grpc.CallOption) (*GetGroupContactApplicationListOut, error) {
	out := new(GetGroupContactApplicationListOut)
	err := c.cc.Invoke(ctx, "/pb.user/getGroupContactApplicationList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateContactApplication(ctx context.Context, in *CreateContactApplicationIn, opts ...grpc.CallOption) (*CreateContactApplicationOut, error) {
	out := new(CreateContactApplicationOut)
	err := c.cc.Invoke(ctx, "/pb.user/createContactApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateContactApplication(ctx context.Context, in *UpdateContactApplicationIn, opts ...grpc.CallOption) (*UpdateContactApplicationOut, error) {
	out := new(UpdateContactApplicationOut)
	err := c.cc.Invoke(ctx, "/pb.user/updateContactApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Login(context.Context, *LoginIn) (*LoginOut, error)
	Register(context.Context, *RegisterIn) (*RegisterOut, error)
	GetUserList(context.Context, *GetListIn) (*GetListOut, error)
	UpdateUser(context.Context, *UpdateIn) (*UpdateOut, error)
	CreateGroup(context.Context, *CreateGroupIn) (*CreateGroupOut, error)
	UpdateGroup(context.Context, *UpdateGroupIn) (*UpdateGroupOut, error)
	DeleteUserGroupMap(context.Context, *DeleteUserGroupMapIn) (*DeleteUserGroupMapOut, error)
	DeleteGroup(context.Context, *DeleteGroupIn) (*DeleteGroupOut, error)
	GetContactList(context.Context, *GetContactListIn) (*GetContactListOut, error)
	DeleteContact(context.Context, *DeleteContactIn) (*DeleteContactOut, error)
	GetContactApplicationList(context.Context, *GetContactApplicationListIn) (*GetContactApplicationListOut, error)
	GetGroupContactApplicationList(context.Context, *GetGroupContactApplicationListIn) (*GetGroupContactApplicationListOut, error)
	CreateContactApplication(context.Context, *CreateContactApplicationIn) (*CreateContactApplicationOut, error)
	UpdateContactApplication(context.Context, *UpdateContactApplicationIn) (*UpdateContactApplicationOut, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Login(context.Context, *LoginIn) (*LoginOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) Register(context.Context, *RegisterIn) (*RegisterOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) GetUserList(context.Context, *GetListIn) (*GetListOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *UpdateIn) (*UpdateOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) CreateGroup(context.Context, *CreateGroupIn) (*CreateGroupOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedUserServer) UpdateGroup(context.Context, *UpdateGroupIn) (*UpdateGroupOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedUserServer) DeleteUserGroupMap(context.Context, *DeleteUserGroupMapIn) (*DeleteUserGroupMapOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserGroupMap not implemented")
}
func (UnimplementedUserServer) DeleteGroup(context.Context, *DeleteGroupIn) (*DeleteGroupOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedUserServer) GetContactList(context.Context, *GetContactListIn) (*GetContactListOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContactList not implemented")
}
func (UnimplementedUserServer) DeleteContact(context.Context, *DeleteContactIn) (*DeleteContactOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContact not implemented")
}
func (UnimplementedUserServer) GetContactApplicationList(context.Context, *GetContactApplicationListIn) (*GetContactApplicationListOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContactApplicationList not implemented")
}
func (UnimplementedUserServer) GetGroupContactApplicationList(context.Context, *GetGroupContactApplicationListIn) (*GetGroupContactApplicationListOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupContactApplicationList not implemented")
}
func (UnimplementedUserServer) CreateContactApplication(context.Context, *CreateContactApplicationIn) (*CreateContactApplicationOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContactApplication not implemented")
}
func (UnimplementedUserServer) UpdateContactApplication(context.Context, *UpdateContactApplicationIn) (*UpdateContactApplicationOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContactApplication not implemented")
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

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/getUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserList(ctx, req.(*GetListIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/updateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/createGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateGroup(ctx, req.(*CreateGroupIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/updateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateGroup(ctx, req.(*UpdateGroupIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUserGroupMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserGroupMapIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUserGroupMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/deleteUserGroupMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUserGroupMap(ctx, req.(*DeleteUserGroupMapIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/deleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteGroup(ctx, req.(*DeleteGroupIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetContactList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactListIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetContactList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/getContactList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetContactList(ctx, req.(*GetContactListIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContactIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/deleteContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteContact(ctx, req.(*DeleteContactIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetContactApplicationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactApplicationListIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetContactApplicationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/getContactApplicationList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetContactApplicationList(ctx, req.(*GetContactApplicationListIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetGroupContactApplicationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupContactApplicationListIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetGroupContactApplicationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/getGroupContactApplicationList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetGroupContactApplicationList(ctx, req.(*GetGroupContactApplicationListIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateContactApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContactApplicationIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateContactApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/createContactApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateContactApplication(ctx, req.(*CreateContactApplicationIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateContactApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContactApplicationIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateContactApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.user/updateContactApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateContactApplication(ctx, req.(*UpdateContactApplicationIn))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.user",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "getUserList",
			Handler:    _User_GetUserList_Handler,
		},
		{
			MethodName: "updateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "createGroup",
			Handler:    _User_CreateGroup_Handler,
		},
		{
			MethodName: "updateGroup",
			Handler:    _User_UpdateGroup_Handler,
		},
		{
			MethodName: "deleteUserGroupMap",
			Handler:    _User_DeleteUserGroupMap_Handler,
		},
		{
			MethodName: "deleteGroup",
			Handler:    _User_DeleteGroup_Handler,
		},
		{
			MethodName: "getContactList",
			Handler:    _User_GetContactList_Handler,
		},
		{
			MethodName: "deleteContact",
			Handler:    _User_DeleteContact_Handler,
		},
		{
			MethodName: "getContactApplicationList",
			Handler:    _User_GetContactApplicationList_Handler,
		},
		{
			MethodName: "getGroupContactApplicationList",
			Handler:    _User_GetGroupContactApplicationList_Handler,
		},
		{
			MethodName: "createContactApplication",
			Handler:    _User_CreateContactApplication_Handler,
		},
		{
			MethodName: "updateContactApplication",
			Handler:    _User_UpdateContactApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/user/rpc/user.proto",
}
