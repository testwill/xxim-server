// Code generated by goctl. DO NOT EDIT.
// Source: gateway.proto

package server

import (
	"context"

	"github.com/cherish-chat/xxim-server/app/gateway/internal/logic"
	"github.com/cherish-chat/xxim-server/app/gateway/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"
)

type GatewayServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedGatewayServiceServer
}

func NewGatewayServiceServer(svcCtx *svc.ServiceContext) *GatewayServiceServer {
	return &GatewayServiceServer{
		svcCtx: svcCtx,
	}
}

// GatewayGetUserConnection 获取用户的连接
func (s *GatewayServiceServer) GatewayGetUserConnection(ctx context.Context, in *pb.GatewayGetUserConnectionReq) (*pb.GatewayGetUserConnectionResp, error) {
	l := logic.NewGatewayGetUserConnectionLogic(ctx, s.svcCtx)
	return l.GatewayGetUserConnection(in)
}

// GatewayBatchGetUserConnection 批量获取用户的连接
func (s *GatewayServiceServer) GatewayBatchGetUserConnection(ctx context.Context, in *pb.GatewayBatchGetUserConnectionReq) (*pb.GatewayBatchGetUserConnectionResp, error) {
	l := logic.NewGatewayBatchGetUserConnectionLogic(ctx, s.svcCtx)
	return l.GatewayBatchGetUserConnection(in)
}

// GatewayGetConnectionByFilter 通过条件获取用户的连接
func (s *GatewayServiceServer) GatewayGetConnectionByFilter(ctx context.Context, in *pb.GatewayGetConnectionByFilterReq) (*pb.GatewayGetConnectionByFilterResp, error) {
	l := logic.NewGatewayGetConnectionByFilterLogic(ctx, s.svcCtx)
	return l.GatewayGetConnectionByFilter(in)
}

// GatewayWriteDataToWs 向用户的连接写入数据
func (s *GatewayServiceServer) GatewayWriteDataToWs(ctx context.Context, in *pb.GatewayWriteDataToWsReq) (*pb.GatewayWriteDataToWsResp, error) {
	l := logic.NewGatewayWriteDataToWsLogic(ctx, s.svcCtx)
	return l.GatewayWriteDataToWs(in)
}

// GatewayKickWs 踢出用户的连接
func (s *GatewayServiceServer) GatewayKickWs(ctx context.Context, in *pb.GatewayKickWsReq) (*pb.GatewayKickWsResp, error) {
	l := logic.NewGatewayKickWsLogic(ctx, s.svcCtx)
	return l.GatewayKickWs(in)
}

// KeepAlive 保持连接
func (s *GatewayServiceServer) GatewayKeepAlive(ctx context.Context, in *pb.GatewayKeepAliveReq) (*pb.GatewayKeepAliveResp, error) {
	l := logic.NewGatewayKeepAliveLogic(ctx, s.svcCtx)
	return l.GatewayKeepAlive(in)
}
