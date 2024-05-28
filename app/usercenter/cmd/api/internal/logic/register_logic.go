package logic

import (
	"context"
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
	// 执行远程调用，注册
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = l.svcCtx.UserCenterRpc.Register(ctx, &usercenter.RegisterReq{
		Phone:   req.Phone,
		Country: req.Country,
	})
	if err != nil {
		l.Logger.Errorf("rpc UserCenterRpc.Register err. err=%+v", err)
		return nil, xerr.NewCommonErrorByGrpcError(err)
	}

	return
}
