// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: apps/message/rpc/message.proto

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
	Message_SaveMsg_FullMethodName            = "/pb.message/saveMsg"
	Message_GetNextMsgList_FullMethodName     = "/pb.message/GetNextMsgList"
	Message_GetPreviousMsgList_FullMethodName = "/pb.message/GetPreviousMsgList"
	Message_DownloadFile_FullMethodName       = "/pb.message/DownloadFile"
	Message_UploadFile_FullMethodName         = "/pb.message/UploadFile"
)

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	SaveMsg(ctx context.Context, in *SaveMsgIn, opts ...grpc.CallOption) (*SaveMsgOut, error)
	GetNextMsgList(ctx context.Context, in *GetNextMsgListIn, opts ...grpc.CallOption) (*GetNextMsgListOut, error)
	GetPreviousMsgList(ctx context.Context, in *GetPreviousMsgListIn, opts ...grpc.CallOption) (*GetPreviousMsgListOut, error)
	DownloadFile(ctx context.Context, in *DownloadFileIn, opts ...grpc.CallOption) (*DownloadFileOut, error)
	UploadFile(ctx context.Context, in *UploadFileIn, opts ...grpc.CallOption) (*UploadFileOut, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) SaveMsg(ctx context.Context, in *SaveMsgIn, opts ...grpc.CallOption) (*SaveMsgOut, error) {
	out := new(SaveMsgOut)
	err := c.cc.Invoke(ctx, Message_SaveMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) GetNextMsgList(ctx context.Context, in *GetNextMsgListIn, opts ...grpc.CallOption) (*GetNextMsgListOut, error) {
	out := new(GetNextMsgListOut)
	err := c.cc.Invoke(ctx, Message_GetNextMsgList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) GetPreviousMsgList(ctx context.Context, in *GetPreviousMsgListIn, opts ...grpc.CallOption) (*GetPreviousMsgListOut, error) {
	out := new(GetPreviousMsgListOut)
	err := c.cc.Invoke(ctx, Message_GetPreviousMsgList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) DownloadFile(ctx context.Context, in *DownloadFileIn, opts ...grpc.CallOption) (*DownloadFileOut, error) {
	out := new(DownloadFileOut)
	err := c.cc.Invoke(ctx, Message_DownloadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) UploadFile(ctx context.Context, in *UploadFileIn, opts ...grpc.CallOption) (*UploadFileOut, error) {
	out := new(UploadFileOut)
	err := c.cc.Invoke(ctx, Message_UploadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	SaveMsg(context.Context, *SaveMsgIn) (*SaveMsgOut, error)
	GetNextMsgList(context.Context, *GetNextMsgListIn) (*GetNextMsgListOut, error)
	GetPreviousMsgList(context.Context, *GetPreviousMsgListIn) (*GetPreviousMsgListOut, error)
	DownloadFile(context.Context, *DownloadFileIn) (*DownloadFileOut, error)
	UploadFile(context.Context, *UploadFileIn) (*UploadFileOut, error)
	mustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) SaveMsg(context.Context, *SaveMsgIn) (*SaveMsgOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMsg not implemented")
}
func (UnimplementedMessageServer) GetNextMsgList(context.Context, *GetNextMsgListIn) (*GetNextMsgListOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNextMsgList not implemented")
}
func (UnimplementedMessageServer) GetPreviousMsgList(context.Context, *GetPreviousMsgListIn) (*GetPreviousMsgListOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreviousMsgList not implemented")
}
func (UnimplementedMessageServer) DownloadFile(context.Context, *DownloadFileIn) (*DownloadFileOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedMessageServer) UploadFile(context.Context, *UploadFileIn) (*UploadFileOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_SaveMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveMsgIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SaveMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SaveMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SaveMsg(ctx, req.(*SaveMsgIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_GetNextMsgList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNextMsgListIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetNextMsgList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_GetNextMsgList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetNextMsgList(ctx, req.(*GetNextMsgListIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_GetPreviousMsgList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreviousMsgListIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetPreviousMsgList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_GetPreviousMsgList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetPreviousMsgList(ctx, req.(*GetPreviousMsgListIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_DownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFileIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).DownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_DownloadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).DownloadFile(ctx, req.(*DownloadFileIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).UploadFile(ctx, req.(*UploadFileIn))
	}
	return interceptor(ctx, in, info, handler)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "saveMsg",
			Handler:    _Message_SaveMsg_Handler,
		},
		{
			MethodName: "GetNextMsgList",
			Handler:    _Message_GetNextMsgList_Handler,
		},
		{
			MethodName: "GetPreviousMsgList",
			Handler:    _Message_GetPreviousMsgList_Handler,
		},
		{
			MethodName: "DownloadFile",
			Handler:    _Message_DownloadFile_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _Message_UploadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/message/rpc/message.proto",
}
