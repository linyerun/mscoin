package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"mscoin/app/usercenter/cmd/rpc/usercenter"
	"mscoin/common/xerr"
	"time"

	"mscoin/app/usercenter/cmd/api/internal/svc"
	"mscoin/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	// 复制rpc请求参数
	loginReq := &usercenter.LoginReq{}
	if err = copier.Copy(loginReq, req); err != nil {
		l.Logger.Errorf("copier copy *types.LoginReq err. err=%+v", err)
		return nil, xerr.NewCommonErrorByCode(xerr.ServerError)
	}

	// rpc请求
	loginResp, err := l.svcCtx.UserCenterRpc.Login(ctx, loginReq)
	if err != nil {
		l.Logger.Errorf("rpc UserCenterRpc.Login err. err=%+v", err)
		return nil, xerr.NewCommonErrorByGrpcError(err)
	}

	// 复制rpc响应参数
	resp = &types.LoginResp{}
	if err = copier.Copy(resp, loginResp); err != nil {
		l.Logger.Errorf("copier copy *usercenter.LoginResp err. err=%+v", err)
		return nil, xerr.NewCommonErrorByCode(xerr.ServerError)
	}

	return
}
