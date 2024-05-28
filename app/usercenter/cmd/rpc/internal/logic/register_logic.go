package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"mscoin/app/usercenter/model"
	"mscoin/common/constant"
	"mscoin/common/tool"
	"mscoin/common/xerr"

	"mscoin/app/usercenter/cmd/rpc/internal/svc"
	"mscoin/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 人机校验(不需要这个功能)
	//ok := util.CaptchaVerify(in.Captcha.Server, l.svcCtx.Config.Captcha.Vid, l.svcCtx.Config.Captcha.Key, in.Captcha.Token, 2, in.Ip)
	//if !ok {
	//	return nil, errors.Wrap(xerr.NewCommonErrorByCode(xerr.RobotCaptchaVerifyError), "验证码验证失败")
	//}

	// 手机验证码校验
	var code string
	err := l.svcCtx.Cache.Get(constant.RegisterCodeCacheKey+in.Phone, &code)
	if err != nil {
		return nil, errors.Wrap(xerr.NewCommonErrorByCode(xerr.RedisGetKeyExpiredError), "根据key获取redis数据失败")
	}
	if code != in.Code {
		return nil, errors.Wrap(xerr.NewCommonErrorByCode(xerr.UserPhoneCodeError), "用户验证码错误")
	}

	// 判断手机是否被注册
	err = l.svcCtx.Db.Table(model.MemberTableName).Where("mobile_phone=?", in.Phone).Limit(1).Take(&model.Member{}).Error
	if err != gorm.ErrRecordNotFound { // 说明手机号已被注册
		return nil, errors.Wrap(xerr.NewCommonErrorByCode(xerr.MobilePhoneExists), "手机号已被注册")
	}

	// 生成member模型, 存入数据库
	mem := &model.Member{}
	err = tool.DataColumnZero(mem)
	if err != nil {
		return nil, errors.Wrap(xerr.NewCommonErrorByCode(xerr.ServerError), "tool.DataColumnZero初始化零值异常")
	}
	salt, pwd := tool.Encode(in.Password, nil)
	mem.Username = in.Username
	mem.Country = in.Country
	mem.Salt = salt
	mem.Password = pwd
	mem.MobilePhone = in.Phone
	mem.FillSuperPartner(in.SuperPartner)
	mem.PromotionCode = in.Promotion
	mem.MemberLevel = model.General
	mem.Avatar = "https://mszlu.oss-cn-beijing.aliyuncs.com/mscoin/defaultavatar.png"
	err = l.svcCtx.Db.Save(mem).Error
	if err != nil {
		return nil, errors.Wrap(xerr.NewCommonErrorByCode(xerr.RegisterError), "保存member信息入数据库失败")
	}

	return &pb.RegisterResp{}, nil
}
