// Code generated by goctl. DO NOT EDIT.
// Source: message.proto

package server

import (
	"context"

	"github.com/cherish-chat/xxim-server/app/message/internal/logic/messageservice"
	"github.com/cherish-chat/xxim-server/app/message/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"
)

type MessageServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedMessageServiceServer
}

func NewMessageServiceServer(svcCtx *svc.ServiceContext) *MessageServiceServer {
	return &MessageServiceServer{
		svcCtx: svcCtx,
	}
}

// MessageInsert 插入消息
func (s *MessageServiceServer) MessageInsert(ctx context.Context, in *pb.MessageInsertReq) (*pb.MessageInsertResp, error) {
	l := messageservicelogic.NewMessageInsertLogic(ctx, s.svcCtx)
	return l.MessageInsert(in)
}

// MessageSend 发送消息
func (s *MessageServiceServer) MessageSend(ctx context.Context, in *pb.MessageSendReq) (*pb.MessageSendResp, error) {
	l := messageservicelogic.NewMessageSendLogic(ctx, s.svcCtx)
	return l.MessageSend(in)
}

// MessageBatchSend 批量发送消息
func (s *MessageServiceServer) MessageBatchSend(ctx context.Context, in *pb.MessageBatchSendReq) (*pb.MessageBatchSendResp, error) {
	l := messageservicelogic.NewMessageBatchSendLogic(ctx, s.svcCtx)
	return l.MessageBatchSend(in)
}

// MessagePush 推送消息
func (s *MessageServiceServer) MessagePush(ctx context.Context, in *pb.MessagePushReq) (*pb.MessagePushResp, error) {
	l := messageservicelogic.NewMessagePushLogic(ctx, s.svcCtx)
	return l.MessagePush(in)
}
