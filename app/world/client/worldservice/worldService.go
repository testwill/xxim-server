// Code generated by goctl. DO NOT EDIT.
// Source: world.proto

package worldservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	WorldPostSubmitReq  = pb.WorldPostSubmitReq
	WorldPostSubmitResp = pb.WorldPostSubmitResp

	WorldService interface {
		// WorldPostSubmit 世界圈帖子发布
		WorldPostSubmit(ctx context.Context, in *WorldPostSubmitReq, opts ...grpc.CallOption) (*WorldPostSubmitResp, error)
	}

	defaultWorldService struct {
		cli zrpc.Client
	}
)

func NewWorldService(cli zrpc.Client) WorldService {
	return &defaultWorldService{
		cli: cli,
	}
}

// WorldPostSubmit 世界圈帖子发布
func (m *defaultWorldService) WorldPostSubmit(ctx context.Context, in *WorldPostSubmitReq, opts ...grpc.CallOption) (*WorldPostSubmitResp, error) {
	client := pb.NewWorldServiceClient(m.cli.Conn())
	return client.WorldPostSubmit(ctx, in, opts...)
}
