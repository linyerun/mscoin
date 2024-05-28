package httpinterceptor

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"mscoin/common/xerr"
	"net/http"
)

func ErrorInterceptor(err error) (int, any) {
	switch e := err.(type) {
	case xerr.Error:
		// 做出永远返回200, json格式返回值
		return http.StatusOK, map[string]any{"code": e.GetCode(), "msg": e.GetMsg()}
	default:
		// 如果可以转成xerr.Error就转
		if _e, ok := errors.Cause(err).(xerr.Error); ok {
			return http.StatusOK, map[string]any{"code": _e.GetCode(), "msg": _e.GetMsg()}
		}

		// 记录系统异常日志
		logx.Error(err)

		// 返回这种超出我们能处理范围的异常
		return http.StatusOK, map[string]any{"code": xerr.ServerError, "msg": err.Error()}
	}
}
