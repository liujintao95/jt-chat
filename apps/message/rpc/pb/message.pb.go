// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.3
// source: apps/message/rpc/message.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId         string `protobuf:"bytes,1,opt,name=MsgId,proto3" json:"MsgId,omitempty"`
	TransportType int64  `protobuf:"varint,2,opt,name=TransportType,proto3" json:"TransportType,omitempty"`
	From          string `protobuf:"bytes,3,opt,name=From,proto3" json:"From,omitempty"`
	To            string `protobuf:"bytes,4,opt,name=To,proto3" json:"To,omitempty"`
	ToType        int64  `protobuf:"varint,5,opt,name=ToType,proto3" json:"ToType,omitempty"`
	Content       string `protobuf:"bytes,6,opt,name=Content,proto3" json:"Content,omitempty"`
	ContentType   int64  `protobuf:"varint,7,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	FileSuffix    string `protobuf:"bytes,8,opt,name=FileSuffix,proto3" json:"FileSuffix,omitempty"`
	FilePath      string `protobuf:"bytes,9,opt,name=FilePath,proto3" json:"FilePath,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_message_rpc_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_apps_message_rpc_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_apps_message_rpc_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *Message) GetTransportType() int64 {
	if x != nil {
		return x.TransportType
	}
	return 0
}

func (x *Message) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Message) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Message) GetToType() int64 {
	if x != nil {
		return x.ToType
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetContentType() int64 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *Message) GetFileSuffix() string {
	if x != nil {
		return x.FileSuffix
	}
	return ""
}

func (x *Message) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type CreateMsgIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *Message `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateMsgIn) Reset() {
	*x = CreateMsgIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_message_rpc_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMsgIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMsgIn) ProtoMessage() {}

func (x *CreateMsgIn) ProtoReflect() protoreflect.Message {
	mi := &file_apps_message_rpc_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMsgIn.ProtoReflect.Descriptor instead.
func (*CreateMsgIn) Descriptor() ([]byte, []int) {
	return file_apps_message_rpc_message_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMsgIn) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

type CreateMsgOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateMsgOut) Reset() {
	*x = CreateMsgOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_message_rpc_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMsgOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMsgOut) ProtoMessage() {}

func (x *CreateMsgOut) ProtoReflect() protoreflect.Message {
	mi := &file_apps_message_rpc_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMsgOut.ProtoReflect.Descriptor instead.
func (*CreateMsgOut) Descriptor() ([]byte, []int) {
	return file_apps_message_rpc_message_proto_rawDescGZIP(), []int{2}
}

type GetNextMsgListIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId    string `protobuf:"bytes,1,opt,name=MsgId,proto3" json:"MsgId,omitempty"`
	TargetId string `protobuf:"bytes,2,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	Size     int64  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (x *GetNextMsgListIn) Reset() {
	*x = GetNextMsgListIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_message_rpc_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNextMsgListIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNextMsgListIn) ProtoMessage() {}

func (x *GetNextMsgListIn) ProtoReflect() protoreflect.Message {
	mi := &file_apps_message_rpc_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNextMsgListIn.ProtoReflect.Descriptor instead.
func (*GetNextMsgListIn) Descriptor() ([]byte, []int) {
	return file_apps_message_rpc_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetNextMsgListIn) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *GetNextMsgListIn) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *GetNextMsgListIn) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type GetNextMsgListOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageList []*Message `protobuf:"bytes,1,rep,name=MessageList,proto3" json:"MessageList,omitempty"`
}

func (x *GetNextMsgListOut) Reset() {
	*x = GetNextMsgListOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_message_rpc_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNextMsgListOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNextMsgListOut) ProtoMessage() {}

func (x *GetNextMsgListOut) ProtoReflect() protoreflect.Message {
	mi := &file_apps_message_rpc_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNextMsgListOut.ProtoReflect.Descriptor instead.
func (*GetNextMsgListOut) Descriptor() ([]byte, []int) {
	return file_apps_message_rpc_message_proto_rawDescGZIP(), []int{4}
}

func (x *GetNextMsgListOut) GetMessageList() []*Message {
	if x != nil {
		return x.MessageList
	}
	return nil
}

type GetPreviousMsgListIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId    string `protobuf:"bytes,1,opt,name=MsgId,proto3" json:"MsgId,omitempty"`
	TargetId string `protobuf:"bytes,2,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	Size     int64  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (x *GetPreviousMsgListIn) Reset() {
	*x = GetPreviousMsgListIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_message_rpc_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPreviousMsgListIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPreviousMsgListIn) ProtoMessage() {}

func (x *GetPreviousMsgListIn) ProtoReflect() protoreflect.Message {
	mi := &file_apps_message_rpc_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPreviousMsgListIn.ProtoReflect.Descriptor instead.
func (*GetPreviousMsgListIn) Descriptor() ([]byte, []int) {
	return file_apps_message_rpc_message_proto_rawDescGZIP(), []int{5}
}

func (x *GetPreviousMsgListIn) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *GetPreviousMsgListIn) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *GetPreviousMsgListIn) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type GetPreviousMsgListOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageList []*Message `protobuf:"bytes,1,rep,name=MessageList,proto3" json:"MessageList,omitempty"`
}

func (x *GetPreviousMsgListOut) Reset() {
	*x = GetPreviousMsgListOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_message_rpc_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPreviousMsgListOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPreviousMsgListOut) ProtoMessage() {}

func (x *GetPreviousMsgListOut) ProtoReflect() protoreflect.Message {
	mi := &file_apps_message_rpc_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPreviousMsgListOut.ProtoReflect.Descriptor instead.
func (*GetPreviousMsgListOut) Descriptor() ([]byte, []int) {
	return file_apps_message_rpc_message_proto_rawDescGZIP(), []int{6}
}

func (x *GetPreviousMsgListOut) GetMessageList() []*Message {
	if x != nil {
		return x.MessageList
	}
	return nil
}

type DownloadFileIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=FilePath,proto3" json:"FilePath,omitempty"`
}

func (x *DownloadFileIn) Reset() {
	*x = DownloadFileIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_message_rpc_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileIn) ProtoMessage() {}

func (x *DownloadFileIn) ProtoReflect() protoreflect.Message {
	mi := &file_apps_message_rpc_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileIn.ProtoReflect.Descriptor instead.
func (*DownloadFileIn) Descriptor() ([]byte, []int) {
	return file_apps_message_rpc_message_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadFileIn) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type DownloadFileOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *DownloadFileOut) Reset() {
	*x = DownloadFileOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_message_rpc_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileOut) ProtoMessage() {}

func (x *DownloadFileOut) ProtoReflect() protoreflect.Message {
	mi := &file_apps_message_rpc_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileOut.ProtoReflect.Descriptor instead.
func (*DownloadFileOut) Descriptor() ([]byte, []int) {
	return file_apps_message_rpc_message_proto_rawDescGZIP(), []int{8}
}

func (x *DownloadFileOut) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type UploadFileIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *UploadFileIn) Reset() {
	*x = UploadFileIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_message_rpc_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileIn) ProtoMessage() {}

func (x *UploadFileIn) ProtoReflect() protoreflect.Message {
	mi := &file_apps_message_rpc_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileIn.ProtoReflect.Descriptor instead.
func (*UploadFileIn) Descriptor() ([]byte, []int) {
	return file_apps_message_rpc_message_proto_rawDescGZIP(), []int{9}
}

func (x *UploadFileIn) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type UploadFileOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=FilePath,proto3" json:"FilePath,omitempty"`
}

func (x *UploadFileOut) Reset() {
	*x = UploadFileOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_message_rpc_message_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileOut) ProtoMessage() {}

func (x *UploadFileOut) ProtoReflect() protoreflect.Message {
	mi := &file_apps_message_rpc_message_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileOut.ProtoReflect.Descriptor instead.
func (*UploadFileOut) Descriptor() ([]byte, []int) {
	return file_apps_message_rpc_message_proto_rawDescGZIP(), []int{10}
}

func (x *UploadFileOut) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

var File_apps_message_rpc_message_proto protoreflect.FileDescriptor

var file_apps_message_rpc_message_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0xf9, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x46, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x54, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x54, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x75, 0x66, 0x66,
	0x69, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x75,
	0x66, 0x66, 0x69, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x2c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x12,
	0x1d, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x0e,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x4f, 0x75, 0x74, 0x22, 0x58,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x42, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x78, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x2d, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5c, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x46, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x75, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x25, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x4f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x32, 0xaf, 0x02, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x73,
	0x67, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67,
	0x49, 0x6e, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x73,
	0x67, 0x4f, 0x75, 0x74, 0x12, 0x3d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x4d,
	0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x78, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x75, 0x74, 0x12, 0x49, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x37,
	0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_message_rpc_message_proto_rawDescOnce sync.Once
	file_apps_message_rpc_message_proto_rawDescData = file_apps_message_rpc_message_proto_rawDesc
)

func file_apps_message_rpc_message_proto_rawDescGZIP() []byte {
	file_apps_message_rpc_message_proto_rawDescOnce.Do(func() {
		file_apps_message_rpc_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_message_rpc_message_proto_rawDescData)
	})
	return file_apps_message_rpc_message_proto_rawDescData
}

var file_apps_message_rpc_message_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_apps_message_rpc_message_proto_goTypes = []interface{}{
	(*Message)(nil),               // 0: pb.Message
	(*CreateMsgIn)(nil),           // 1: pb.CreateMsgIn
	(*CreateMsgOut)(nil),          // 2: pb.CreateMsgOut
	(*GetNextMsgListIn)(nil),      // 3: pb.GetNextMsgListIn
	(*GetNextMsgListOut)(nil),     // 4: pb.GetNextMsgListOut
	(*GetPreviousMsgListIn)(nil),  // 5: pb.GetPreviousMsgListIn
	(*GetPreviousMsgListOut)(nil), // 6: pb.GetPreviousMsgListOut
	(*DownloadFileIn)(nil),        // 7: pb.DownloadFileIn
	(*DownloadFileOut)(nil),       // 8: pb.DownloadFileOut
	(*UploadFileIn)(nil),          // 9: pb.UploadFileIn
	(*UploadFileOut)(nil),         // 10: pb.UploadFileOut
}
var file_apps_message_rpc_message_proto_depIdxs = []int32{
	0,  // 0: pb.CreateMsgIn.msg:type_name -> pb.Message
	0,  // 1: pb.GetNextMsgListOut.MessageList:type_name -> pb.Message
	0,  // 2: pb.GetPreviousMsgListOut.MessageList:type_name -> pb.Message
	1,  // 3: pb.message.CreateMsg:input_type -> pb.CreateMsgIn
	3,  // 4: pb.message.GetNextMsgList:input_type -> pb.GetNextMsgListIn
	5,  // 5: pb.message.GetPreviousMsgList:input_type -> pb.GetPreviousMsgListIn
	7,  // 6: pb.message.DownloadFile:input_type -> pb.DownloadFileIn
	9,  // 7: pb.message.UploadFile:input_type -> pb.UploadFileIn
	2,  // 8: pb.message.CreateMsg:output_type -> pb.CreateMsgOut
	4,  // 9: pb.message.GetNextMsgList:output_type -> pb.GetNextMsgListOut
	6,  // 10: pb.message.GetPreviousMsgList:output_type -> pb.GetPreviousMsgListOut
	8,  // 11: pb.message.DownloadFile:output_type -> pb.DownloadFileOut
	10, // 12: pb.message.UploadFile:output_type -> pb.UploadFileOut
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_apps_message_rpc_message_proto_init() }
func file_apps_message_rpc_message_proto_init() {
	if File_apps_message_rpc_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_message_rpc_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_message_rpc_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMsgIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_message_rpc_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMsgOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_message_rpc_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNextMsgListIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_message_rpc_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNextMsgListOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_message_rpc_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPreviousMsgListIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_message_rpc_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPreviousMsgListOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_message_rpc_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_message_rpc_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_message_rpc_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_message_rpc_message_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_message_rpc_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_message_rpc_message_proto_goTypes,
		DependencyIndexes: file_apps_message_rpc_message_proto_depIdxs,
		MessageInfos:      file_apps_message_rpc_message_proto_msgTypes,
	}.Build()
	File_apps_message_rpc_message_proto = out.File
	file_apps_message_rpc_message_proto_rawDesc = nil
	file_apps_message_rpc_message_proto_goTypes = nil
	file_apps_message_rpc_message_proto_depIdxs = nil
}