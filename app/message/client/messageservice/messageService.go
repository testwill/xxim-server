// Code generated by goctl. DO NOT EDIT.
// Source: message.proto

package messageservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ListNoticeReq                 = pb.ListNoticeReq
	ListNoticeResp                = pb.ListNoticeResp
	Message                       = pb.Message
	MessageBatchSendReq           = pb.MessageBatchSendReq
	MessageBatchSendResp          = pb.MessageBatchSendResp
	MessageContentText            = pb.MessageContentText
	MessageContentText_Item       = pb.MessageContentText_Item
	MessageContentText_Item_At    = pb.MessageContentText_Item_At
	MessageContentText_Item_Image = pb.MessageContentText_Item_Image
	MessageInsertReq              = pb.MessageInsertReq
	MessageInsertResp             = pb.MessageInsertResp
	MessagePushReq                = pb.MessagePushReq
	MessagePushResp               = pb.MessagePushResp
	MessageSendReq                = pb.MessageSendReq
	MessageSendResp               = pb.MessageSendResp
	Message_Option                = pb.Message_Option
	Message_Sender                = pb.Message_Sender
	Notice                        = pb.Notice
	NoticeBatchSendReq            = pb.NoticeBatchSendReq
	NoticeBatchSendResp           = pb.NoticeBatchSendResp
	NoticeContentNewFriendRequest = pb.NoticeContentNewFriendRequest
	NoticeSendReq                 = pb.NoticeSendReq
	NoticeSendResp                = pb.NoticeSendResp

	MessageService interface {
		// MessageInsert 插入消息
		MessageInsert(ctx context.Context, in *MessageInsertReq, opts ...grpc.CallOption) (*MessageInsertResp, error)
		// MessageSend 发送消息
		MessageSend(ctx context.Context, in *MessageSendReq, opts ...grpc.CallOption) (*MessageSendResp, error)
		// MessageBatchSend 批量发送消息
		MessageBatchSend(ctx context.Context, in *MessageBatchSendReq, opts ...grpc.CallOption) (*MessageBatchSendResp, error)
		// MessagePush 推送消息
		MessagePush(ctx context.Context, in *MessagePushReq, opts ...grpc.CallOption) (*MessagePushResp, error)
	}

	defaultMessageService struct {
		cli zrpc.Client
	}
)

func NewMessageService(cli zrpc.Client) MessageService {
	return &defaultMessageService{
		cli: cli,
	}
}

// MessageInsert 插入消息
func (m *defaultMessageService) MessageInsert(ctx context.Context, in *MessageInsertReq, opts ...grpc.CallOption) (*MessageInsertResp, error) {
	client := pb.NewMessageServiceClient(m.cli.Conn())
	return client.MessageInsert(ctx, in, opts...)
}

// MessageSend 发送消息
func (m *defaultMessageService) MessageSend(ctx context.Context, in *MessageSendReq, opts ...grpc.CallOption) (*MessageSendResp, error) {
	client := pb.NewMessageServiceClient(m.cli.Conn())
	return client.MessageSend(ctx, in, opts...)
}

// MessageBatchSend 批量发送消息
func (m *defaultMessageService) MessageBatchSend(ctx context.Context, in *MessageBatchSendReq, opts ...grpc.CallOption) (*MessageBatchSendResp, error) {
	client := pb.NewMessageServiceClient(m.cli.Conn())
	return client.MessageBatchSend(ctx, in, opts...)
}

// MessagePush 推送消息
func (m *defaultMessageService) MessagePush(ctx context.Context, in *MessagePushReq, opts ...grpc.CallOption) (*MessagePushResp, error) {
	client := pb.NewMessageServiceClient(m.cli.Conn())
	return client.MessagePush(ctx, in, opts...)
}
