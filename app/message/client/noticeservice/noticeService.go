// Code generated by goctl. DO NOT EDIT.
// Source: message.proto

package noticeservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ListNoticeReq                 = pb.ListNoticeReq
	ListNoticeResp                = pb.ListNoticeResp
	MessageInsertReq              = pb.MessageInsertReq
	MessageInsertResp             = pb.MessageInsertResp
	Notice                        = pb.Notice
	NoticeBatchSendReq            = pb.NoticeBatchSendReq
	NoticeBatchSendResp           = pb.NoticeBatchSendResp
	NoticeContentNewFriendRequest = pb.NoticeContentNewFriendRequest
	NoticeSendReq                 = pb.NoticeSendReq
	NoticeSendResp                = pb.NoticeSendResp

	NoticeService interface {
		// NoticeSend 通知发送
		NoticeSend(ctx context.Context, in *NoticeSendReq, opts ...grpc.CallOption) (*NoticeSendResp, error)
		// NoticeBatchSend 通知批量发送
		NoticeBatchSend(ctx context.Context, in *NoticeBatchSendReq, opts ...grpc.CallOption) (*NoticeBatchSendResp, error)
		// ListNotice 获取通知列表
		ListNotice(ctx context.Context, in *ListNoticeReq, opts ...grpc.CallOption) (*ListNoticeResp, error)
	}

	defaultNoticeService struct {
		cli zrpc.Client
	}
)

func NewNoticeService(cli zrpc.Client) NoticeService {
	return &defaultNoticeService{
		cli: cli,
	}
}

// NoticeSend 通知发送
func (m *defaultNoticeService) NoticeSend(ctx context.Context, in *NoticeSendReq, opts ...grpc.CallOption) (*NoticeSendResp, error) {
	client := pb.NewNoticeServiceClient(m.cli.Conn())
	return client.NoticeSend(ctx, in, opts...)
}

// NoticeBatchSend 通知批量发送
func (m *defaultNoticeService) NoticeBatchSend(ctx context.Context, in *NoticeBatchSendReq, opts ...grpc.CallOption) (*NoticeBatchSendResp, error) {
	client := pb.NewNoticeServiceClient(m.cli.Conn())
	return client.NoticeBatchSend(ctx, in, opts...)
}

// ListNotice 获取通知列表
func (m *defaultNoticeService) ListNotice(ctx context.Context, in *ListNoticeReq, opts ...grpc.CallOption) (*ListNoticeResp, error) {
	client := pb.NewNoticeServiceClient(m.cli.Conn())
	return client.ListNotice(ctx, in, opts...)
}
