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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 复制请求参数
	registerReq := &usercenter.RegisterReq{}
	if err = copier.Copy(registerReq, req); err != nil {
		l.Logger.Errorf("copier copy *types.RegisterReq err. err=%+v", err)
		return nil, xerr.NewCommonErrorByCode(xerr.ServerError)
	}

	// 执行远程调用，注册
	_, err = l.svcCtx.UserCenterRpc.Register(ctx, registerReq)
	if err != nil {
		l.Logger.Errorf("rpc UserCenterRpc.Register err. err=%+v", err)
		return nil, xerr.NewCommonErrorByGrpcError(err)
	}

	// 没远程调用异常就是调用注册成功

	return
}
