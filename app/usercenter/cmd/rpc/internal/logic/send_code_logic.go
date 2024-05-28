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

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCodeLogic) SendCode(in *pb.SendCodeReq) (*pb.SendCodeResp, error) {
	// 获取code
	code := tool.RandCode(6)

	// 假设调用短信平台发送验证码

	// 打印验证码
	l.Infof("phone=%s, country=%s. 验证码为%s", in.Phone, in.Country, code)

	// 缓存验证码(多加一个60秒随机是避免短时间内多key过期)
	err := l.svcCtx.Cache.SetWithExpire(constant.RegisterCodeCacheKey+in.Phone, code, 15*time.Minute+time.Second*time.Duration(rand.Intn(60)))
	if err != nil {
		return nil, errors.Wrap(xerr.NewCommonErrorByCode(xerr.RedisSetKeyExpiredError), err.Error())
	}

	return &pb.SendCodeResp{}, nil
}
