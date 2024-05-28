package logic

import (
	"context"
	"github.com/pkg/errors"
	"math/rand"
	"mscoin/common/constant"
	"mscoin/common/tool"
	"mscoin/common/xerr"
	"time"

	"mscoin/app/usercenter/cmd/rpc/internal/svc"
	"mscoin/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawSendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWithdrawSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawSendCodeLogic {
	return &WithdrawSendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WithdrawSendCodeLogic) WithdrawSendCode(in *pb.WithdrawSendCodeReq) (*pb.WithdrawSendCodeResp, error) {
	// 获取code
	code := tool.RandCode(6)

	// 假设调用短信平台发送验证码

	// 打印验证码
	l.Infof("phone=%s. Withdraw验证码为%s", in.Phone, code)

	// 缓存验证码(多加一个60秒随机是避免短时间内多key过期)
	err := l.svcCtx.Cache.SetWithExpire(constant.WithdrawCodeCacheKey+in.Phone, code, 5*time.Minute+time.Second*time.Duration(rand.Intn(60)))
	if err != nil {
		return nil, errors.Wrap(xerr.NewCommonErrorByCode(xerr.RedisSetKeyExpiredError), err.Error())
	}

	return &pb.WithdrawSendCodeResp{}, nil
}
