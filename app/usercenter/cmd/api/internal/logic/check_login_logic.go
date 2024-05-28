package logic

import (
	"context"
	"mscoin/common/tool"
	"mscoin/common/xerr"

	"mscoin/app/usercenter/cmd/api/internal/svc"
	"mscoin/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLoginLogic {
	return &CheckLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLoginLogic) CheckLogin(req *types.CheckLoginReq) (resp *types.CheckLoginResp, err error) {
	userId, err := tool.ParseToken(req.Token, l.svcCtx.Config.JWT.AccessSecret)
	if err == tool.TokenInValidError {
		return nil, xerr.NewCommonErrorByCode(xerr.TokenInValidError)
	}

	return &types.CheckLoginResp{IsValid: userId > 0}, nil
}
