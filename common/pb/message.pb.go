// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: message.proto

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

type MessageInsertReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
}

func (x *MessageInsertReq) Reset() {
	*x = MessageInsertReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageInsertReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageInsertReq) ProtoMessage() {}

func (x *MessageInsertReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageInsertReq.ProtoReflect.Descriptor instead.
func (*MessageInsertReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *MessageInsertReq) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type MessageInsertResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
}

func (x *MessageInsertResp) Reset() {
	*x = MessageInsertResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageInsertResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageInsertResp) ProtoMessage() {}

func (x *MessageInsertResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageInsertResp.ProtoReflect.Descriptor instead.
func (*MessageInsertResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *MessageInsertResp) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

//Notice 通知
type Notice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//通知id 因为通知永远都是服务端发送给客户端的，所以通知id一定唯一。
	NoticeId string `protobuf:"bytes,1,opt,name=NoticeId,proto3" json:"NoticeId"`
	//发送到哪个会话
	ConversationId string `protobuf:"bytes,2,opt,name=conversationId,proto3" json:"conversationId"` // 单聊: 那么该值为接受者的id；群聊：那么该值为群id；订阅号：那么该值为订阅号id
	//会话类型
	ConversationType ConversationType `protobuf:"varint,3,opt,name=conversationType,proto3,enum=pb.ConversationType" json:"conversationType"`
	//通知内容
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content"`
	//通知类型
	ContentType NoticeContentType `protobuf:"varint,5,opt,name=contentType,proto3,enum=pb.NoticeContentType" json:"contentType"`
	//通知的更新时间
	UpdateTime int64 `protobuf:"varint,6,opt,name=updateTime,proto3" json:"updateTime"`
}

func (x *Notice) Reset() {
	*x = Notice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notice) ProtoMessage() {}

func (x *Notice) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notice.ProtoReflect.Descriptor instead.
func (*Notice) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *Notice) GetNoticeId() string {
	if x != nil {
		return x.NoticeId
	}
	return ""
}

func (x *Notice) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

func (x *Notice) GetConversationType() ConversationType {
	if x != nil {
		return x.ConversationType
	}
	return ConversationType_Single
}

func (x *Notice) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Notice) GetContentType() NoticeContentType {
	if x != nil {
		return x.ContentType
	}
	return NoticeContentType_NewFriendRequest
}

func (x *Notice) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

//NoticeContentNewFriendRequest 通知内容-新的好友请求
type NoticeContentNewFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoticeContentNewFriendRequest) Reset() {
	*x = NoticeContentNewFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeContentNewFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeContentNewFriendRequest) ProtoMessage() {}

func (x *NoticeContentNewFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeContentNewFriendRequest.ProtoReflect.Descriptor instead.
func (*NoticeContentNewFriendRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

type NoticeSendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
	Notice *Notice        `protobuf:"bytes,2,opt,name=notice,proto3" json:"notice"`
	//追加用户
	UserIds []string `protobuf:"bytes,3,rep,name=userIds,proto3" json:"userIds"`
	//是否广播
	Broadcast bool `protobuf:"varint,4,opt,name=broadcast,proto3" json:"broadcast"`
}

func (x *NoticeSendReq) Reset() {
	*x = NoticeSendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeSendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeSendReq) ProtoMessage() {}

func (x *NoticeSendReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeSendReq.ProtoReflect.Descriptor instead.
func (*NoticeSendReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *NoticeSendReq) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *NoticeSendReq) GetNotice() *Notice {
	if x != nil {
		return x.Notice
	}
	return nil
}

func (x *NoticeSendReq) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *NoticeSendReq) GetBroadcast() bool {
	if x != nil {
		return x.Broadcast
	}
	return false
}

type NoticeSendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
}

func (x *NoticeSendResp) Reset() {
	*x = NoticeSendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeSendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeSendResp) ProtoMessage() {}

func (x *NoticeSendResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeSendResp.ProtoReflect.Descriptor instead.
func (*NoticeSendResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *NoticeSendResp) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type NoticeBatchSendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *RequestHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
	Notices []*NoticeSendReq `protobuf:"bytes,2,rep,name=notices,proto3" json:"notices"`
}

func (x *NoticeBatchSendReq) Reset() {
	*x = NoticeBatchSendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeBatchSendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeBatchSendReq) ProtoMessage() {}

func (x *NoticeBatchSendReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeBatchSendReq.ProtoReflect.Descriptor instead.
func (*NoticeBatchSendReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *NoticeBatchSendReq) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *NoticeBatchSendReq) GetNotices() []*NoticeSendReq {
	if x != nil {
		return x.Notices
	}
	return nil
}

type NoticeBatchSendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
}

func (x *NoticeBatchSendResp) Reset() {
	*x = NoticeBatchSendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeBatchSendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeBatchSendResp) ProtoMessage() {}

func (x *NoticeBatchSendResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeBatchSendResp.ProtoReflect.Descriptor instead.
func (*NoticeBatchSendResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *NoticeBatchSendResp) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type ListNoticeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
	//会话id
	ConversationId string `protobuf:"bytes,2,opt,name=conversationId,proto3" json:"conversationId"`
	//会话类型
	ConversationType ConversationType `protobuf:"varint,3,opt,name=conversationType,proto3,enum=pb.ConversationType" json:"conversationType"`
	//UpdateTimeGt 通知更新时间大于
	UpdateTimeGt int64 `protobuf:"varint,4,opt,name=UpdateTimeGt,proto3" json:"UpdateTimeGt"`
	//Limit 限制条数
	Limit int64 `protobuf:"varint,5,opt,name=Limit,proto3" json:"Limit"`
}

func (x *ListNoticeReq) Reset() {
	*x = ListNoticeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNoticeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNoticeReq) ProtoMessage() {}

func (x *ListNoticeReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNoticeReq.ProtoReflect.Descriptor instead.
func (*ListNoticeReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{8}
}

func (x *ListNoticeReq) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListNoticeReq) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

func (x *ListNoticeReq) GetConversationType() ConversationType {
	if x != nil {
		return x.ConversationType
	}
	return ConversationType_Single
}

func (x *ListNoticeReq) GetUpdateTimeGt() int64 {
	if x != nil {
		return x.UpdateTimeGt
	}
	return 0
}

func (x *ListNoticeReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListNoticeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
	Notices []*Notice       `protobuf:"bytes,2,rep,name=notices,proto3" json:"notices"`
}

func (x *ListNoticeResp) Reset() {
	*x = ListNoticeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNoticeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNoticeResp) ProtoMessage() {}

func (x *ListNoticeResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNoticeResp.ProtoReflect.Descriptor instead.
func (*ListNoticeResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{9}
}

func (x *ListNoticeResp) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListNoticeResp) GetNotices() []*Notice {
	if x != nil {
		return x.Notices
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3d, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x22, 0x3f, 0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x22, 0x81, 0x02, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x40, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4e, 0x65, 0x77, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52,
	0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x22,
	0x3c, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x6c, 0x0a,
	0x12, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b,
	0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x52, 0x07, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x13, 0x4e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0xde,
	0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x29, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x47, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x47, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x62, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x07, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x07, 0x6e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x73, 0x32, 0x4e, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x32, 0xbd, 0x01, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x6e, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x0f, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_message_proto_goTypes = []interface{}{
	(*MessageInsertReq)(nil),              // 0: pb.MessageInsertReq
	(*MessageInsertResp)(nil),             // 1: pb.MessageInsertResp
	(*Notice)(nil),                        // 2: pb.Notice
	(*NoticeContentNewFriendRequest)(nil), // 3: pb.NoticeContentNewFriendRequest
	(*NoticeSendReq)(nil),                 // 4: pb.NoticeSendReq
	(*NoticeSendResp)(nil),                // 5: pb.NoticeSendResp
	(*NoticeBatchSendReq)(nil),            // 6: pb.NoticeBatchSendReq
	(*NoticeBatchSendResp)(nil),           // 7: pb.NoticeBatchSendResp
	(*ListNoticeReq)(nil),                 // 8: pb.ListNoticeReq
	(*ListNoticeResp)(nil),                // 9: pb.ListNoticeResp
	(*RequestHeader)(nil),                 // 10: pb.RequestHeader
	(*ResponseHeader)(nil),                // 11: pb.ResponseHeader
	(ConversationType)(0),                 // 12: pb.ConversationType
	(NoticeContentType)(0),                // 13: pb.NoticeContentType
}
var file_message_proto_depIdxs = []int32{
	10, // 0: pb.MessageInsertReq.header:type_name -> pb.RequestHeader
	11, // 1: pb.MessageInsertResp.header:type_name -> pb.ResponseHeader
	12, // 2: pb.Notice.conversationType:type_name -> pb.ConversationType
	13, // 3: pb.Notice.contentType:type_name -> pb.NoticeContentType
	10, // 4: pb.NoticeSendReq.header:type_name -> pb.RequestHeader
	2,  // 5: pb.NoticeSendReq.notice:type_name -> pb.Notice
	11, // 6: pb.NoticeSendResp.header:type_name -> pb.ResponseHeader
	10, // 7: pb.NoticeBatchSendReq.header:type_name -> pb.RequestHeader
	4,  // 8: pb.NoticeBatchSendReq.notices:type_name -> pb.NoticeSendReq
	11, // 9: pb.NoticeBatchSendResp.header:type_name -> pb.ResponseHeader
	10, // 10: pb.ListNoticeReq.header:type_name -> pb.RequestHeader
	12, // 11: pb.ListNoticeReq.conversationType:type_name -> pb.ConversationType
	11, // 12: pb.ListNoticeResp.header:type_name -> pb.ResponseHeader
	2,  // 13: pb.ListNoticeResp.notices:type_name -> pb.Notice
	0,  // 14: pb.messageService.MessageInsert:input_type -> pb.MessageInsertReq
	4,  // 15: pb.noticeService.NoticeSend:input_type -> pb.NoticeSendReq
	6,  // 16: pb.noticeService.NoticeBatchSend:input_type -> pb.NoticeBatchSendReq
	8,  // 17: pb.noticeService.ListNotice:input_type -> pb.ListNoticeReq
	1,  // 18: pb.messageService.MessageInsert:output_type -> pb.MessageInsertResp
	5,  // 19: pb.noticeService.NoticeSend:output_type -> pb.NoticeSendResp
	7,  // 20: pb.noticeService.NoticeBatchSend:output_type -> pb.NoticeBatchSendResp
	9,  // 21: pb.noticeService.ListNotice:output_type -> pb.ListNoticeResp
	18, // [18:22] is the sub-list for method output_type
	14, // [14:18] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageInsertReq); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageInsertResp); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notice); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeContentNewFriendRequest); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeSendReq); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeSendResp); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeBatchSendReq); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeBatchSendResp); i {
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
		file_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNoticeReq); i {
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
		file_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNoticeResp); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
