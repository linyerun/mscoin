package logic

import (
	"context"
	"mscoin/app/usercenter/cmd/api/internal/svc"
	"mscoin/app/usercenter/cmd/api/internal/types"
	"mscoin/app/usercenter/cmd/rpc/usercenter"
	"mscoin/common/xerr"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeReq) (resp *types.SendCodeResp, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 执行远程调用，发送短信
	_, err = l.svcCtx.UserCenterRpc.SendCode(ctx, &usercenter.SendCodeReq{
		Phone:   req.Phone,
		Country: req.Country,
	})
	if err != nil {
		l.Logger.Errorf("rpc UserCenterRpc.SendCode err. err=%+v", err)
		return nil, xerr.NewCommonErrorByGrpcError(err)
	}

	return
}
