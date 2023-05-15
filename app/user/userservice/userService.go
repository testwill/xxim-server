// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserRegisterReq  = pb.UserRegisterReq
	UserRegisterResp = pb.UserRegisterResp

	UserService interface {
		// UserRegister 用户注册
		UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

// UserRegister 用户注册
func (m *defaultUserService) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.UserRegister(ctx, in, opts...)
}