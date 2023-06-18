// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: message.proto

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

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	//MessageInsert 插入消息
	//二次开发时，需要实现该接口
	MessageInsert(ctx context.Context, in *MessageInsertReq, opts ...grpc.CallOption) (*MessageInsertResp, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) MessageInsert(ctx context.Context, in *MessageInsertReq, opts ...grpc.CallOption) (*MessageInsertResp, error) {
	out := new(MessageInsertResp)
	err := c.cc.Invoke(ctx, "/pb.messageService/MessageInsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	//MessageInsert 插入消息
	//二次开发时，需要实现该接口
	MessageInsert(context.Context, *MessageInsertReq) (*MessageInsertResp, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) MessageInsert(context.Context, *MessageInsertReq) (*MessageInsertResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageInsert not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_MessageInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageInsertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MessageInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.messageService/MessageInsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MessageInsert(ctx, req.(*MessageInsertReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.messageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageInsert",
			Handler:    _MessageService_MessageInsert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

// NoticeServiceClient is the client API for NoticeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoticeServiceClient interface {
	//NoticeSend 通知发送
	NoticeSend(ctx context.Context, in *NoticeSendReq, opts ...grpc.CallOption) (*NoticeSendResp, error)
	//NoticeBatchSend 通知批量发送
	NoticeBatchSend(ctx context.Context, in *NoticeBatchSendReq, opts ...grpc.CallOption) (*NoticeBatchSendResp, error)
	//ListNotice 获取通知列表
	ListNotice(ctx context.Context, in *ListNoticeReq, opts ...grpc.CallOption) (*ListNoticeResp, error)
}

type noticeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoticeServiceClient(cc grpc.ClientConnInterface) NoticeServiceClient {
	return &noticeServiceClient{cc}
}

func (c *noticeServiceClient) NoticeSend(ctx context.Context, in *NoticeSendReq, opts ...grpc.CallOption) (*NoticeSendResp, error) {
	out := new(NoticeSendResp)
	err := c.cc.Invoke(ctx, "/pb.noticeService/NoticeSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeServiceClient) NoticeBatchSend(ctx context.Context, in *NoticeBatchSendReq, opts ...grpc.CallOption) (*NoticeBatchSendResp, error) {
	out := new(NoticeBatchSendResp)
	err := c.cc.Invoke(ctx, "/pb.noticeService/NoticeBatchSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeServiceClient) ListNotice(ctx context.Context, in *ListNoticeReq, opts ...grpc.CallOption) (*ListNoticeResp, error) {
	out := new(ListNoticeResp)
	err := c.cc.Invoke(ctx, "/pb.noticeService/ListNotice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoticeServiceServer is the server API for NoticeService service.
// All implementations must embed UnimplementedNoticeServiceServer
// for forward compatibility
type NoticeServiceServer interface {
	//NoticeSend 通知发送
	NoticeSend(context.Context, *NoticeSendReq) (*NoticeSendResp, error)
	//NoticeBatchSend 通知批量发送
	NoticeBatchSend(context.Context, *NoticeBatchSendReq) (*NoticeBatchSendResp, error)
	//ListNotice 获取通知列表
	ListNotice(context.Context, *ListNoticeReq) (*ListNoticeResp, error)
	mustEmbedUnimplementedNoticeServiceServer()
}

// UnimplementedNoticeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNoticeServiceServer struct {
}

func (UnimplementedNoticeServiceServer) NoticeSend(context.Context, *NoticeSendReq) (*NoticeSendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoticeSend not implemented")
}
func (UnimplementedNoticeServiceServer) NoticeBatchSend(context.Context, *NoticeBatchSendReq) (*NoticeBatchSendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoticeBatchSend not implemented")
}
func (UnimplementedNoticeServiceServer) ListNotice(context.Context, *ListNoticeReq) (*ListNoticeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotice not implemented")
}
func (UnimplementedNoticeServiceServer) mustEmbedUnimplementedNoticeServiceServer() {}

// UnsafeNoticeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoticeServiceServer will
// result in compilation errors.
type UnsafeNoticeServiceServer interface {
	mustEmbedUnimplementedNoticeServiceServer()
}

func RegisterNoticeServiceServer(s grpc.ServiceRegistrar, srv NoticeServiceServer) {
	s.RegisterService(&NoticeService_ServiceDesc, srv)
}

func _NoticeService_NoticeSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoticeSendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).NoticeSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.noticeService/NoticeSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).NoticeSend(ctx, req.(*NoticeSendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoticeService_NoticeBatchSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoticeBatchSendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).NoticeBatchSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.noticeService/NoticeBatchSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).NoticeBatchSend(ctx, req.(*NoticeBatchSendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoticeService_ListNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNoticeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).ListNotice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.noticeService/ListNotice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).ListNotice(ctx, req.(*ListNoticeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NoticeService_ServiceDesc is the grpc.ServiceDesc for NoticeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoticeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.noticeService",
	HandlerType: (*NoticeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NoticeSend",
			Handler:    _NoticeService_NoticeSend_Handler,
		},
		{
			MethodName: "NoticeBatchSend",
			Handler:    _NoticeService_NoticeBatchSend_Handler,
		},
		{
			MethodName: "ListNotice",
			Handler:    _NoticeService_ListNotice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
