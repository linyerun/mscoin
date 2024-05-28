// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package usercenter

import (
	"context"

	"mscoin/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CaptchaReq         = pb.CaptchaReq
	FindMemberByIdReq  = pb.FindMemberByIdReq
	FindMemberByIdResp = pb.FindMemberByIdResp
	LoginReq           = pb.LoginReq
	LoginResp          = pb.LoginResp
	RegisterReq        = pb.RegisterReq
	RegisterResp       = pb.RegisterResp
	SendCodeReq        = pb.SendCodeReq
	SendCodeResp       = pb.SendCodeResp

	Usercenter interface {
		SendCode(ctx context.Context, in *SendCodeReq, opts ...grpc.CallOption) (*SendCodeResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		FindMemberById(ctx context.Context, in *FindMemberByIdReq, opts ...grpc.CallOption) (*FindMemberByIdResp, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

func (m *defaultUsercenter) SendCode(ctx context.Context, in *SendCodeReq, opts ...grpc.CallOption) (*SendCodeResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SendCode(ctx, in, opts...)
}

func (m *defaultUsercenter) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUsercenter) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUsercenter) FindMemberById(ctx context.Context, in *FindMemberByIdReq, opts ...grpc.CallOption) (*FindMemberByIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.FindMemberById(ctx, in, opts...)
}
