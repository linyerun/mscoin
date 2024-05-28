package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"mscoin/app/usercenter/cmd/rpc/internal/model"
	"mscoin/common/tool"
	"mscoin/common/xerr"
	"time"

	"mscoin/app/usercenter/cmd/rpc/internal/svc"
	"mscoin/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// 人机验证码校验(忽略)

	// 校验密码
	member := new(model.Member)
	err := l.svcCtx.Db.Table(model.MemberTableName).Where("mobile_phone=?", in.Username).Limit(1).Take(member).Error
	if err == gorm.ErrRecordNotFound { // 密码错误
		return nil, errors.Wrap(xerr.NewCommonErrorByCode(xerr.LoginError), "用户不存在")
	}
	if ok := tool.Verify(in.Password, member.Salt, member.Password, nil); !ok {
		return nil, errors.Wrap(xerr.NewCommonError(xerr.LoginError, "密码错误"), "用户密码错误")
	}

	// 生成token
	token, err := l.getJwtToken(l.svcCtx.Config.JWT.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JWT.AccessExpired, member.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewCommonError(xerr.ServerError, "系统异常,登录失败"), "token generate err: %+v", err)
	}
	loginCount := member.LoginCount + 1

	// 异步执行更新login_count操作
	go func() {
		err = l.svcCtx.Db.Table(model.MemberTableName).Update("login_count", gorm.Expr("login_count + 1")).Error
		if err != nil {
			l.Logger.Errorf("更新id=%d的Member的login_count数据失败, err=%v", member.Id, err)
		}
	}()

	// 返回信息
	return &pb.LoginResp{
		Token:         token,
		Id:            member.Id,
		Username:      member.Username,
		MemberLevel:   member.MemberLevelStr(),
		MemberRate:    member.MemberRate(),
		RealName:      member.RealName,
		Country:       member.Country,
		Avatar:        member.Avatar,
		PromotionCode: member.PromotionCode,
		SuperPartner:  member.SuperPartner,
		LoginCount:    int32(loginCount),
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
