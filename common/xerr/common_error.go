package xerr

import (
	"fmt"
	"google.golang.org/grpc/status"
)

const (
	RobotCaptchaVerifyError = 450
	UserPhoneCodeError      = 451
	MobilePhoneExists       = 452
	LoginError              = 453
	FindInfoError           = 454
	TokenInValidError       = 455

	ServerError = 500

	RedisSetKeyExpiredError = 540
	RedisGetKeyExpiredError = 541
	RegisterError           = 542
)

var codeToMsg = map[int]string{
	ServerError:             "服务异常",
	RedisSetKeyExpiredError: "redis存储有过期时间的键值对异常",
	RedisGetKeyExpiredError: "redis获取有过期时间的键值对异常",
	RobotCaptchaVerifyError: "机器人验证码验证失败",
	UserPhoneCodeError:      "手机验证码错误",
	MobilePhoneExists:       "手机号已存在",
	RegisterError:           "注册失败",
	LoginError:              "登录失败",
	FindInfoError:           "信息不存在",
	TokenInValidError:       "token失效",
}

type Error interface {
	GetCode() int
	GetMsg() string
	Error() string
}

type commonError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCommonError(code int, msg string) Error {
	return &commonError{
		Code: code,
		Msg:  msg,
	}
}

func NewCommonErrorByCode(code int) Error {
	msg, ok := codeToMsg[code]
	if !ok {
		msg = "未知错误"
	}
	return &commonError{Code: code, Msg: msg}
}

func NewCommonErrorByGrpcError(e error) Error {
	if err, ok := status.FromError(e); !ok {
		return NewCommonError(int(err.Code()), err.Message())
	}
	return &commonError{Code: ServerError, Msg: e.Error()}
}

func (c *commonError) GetCode() int {
	return c.Code
}

func (c *commonError) GetMsg() string {
	return c.Msg
}

func (c *commonError) Error() string {
	return fmt.Sprintf("code=%d,msg=%s.", c.Code, c.Msg)
}
