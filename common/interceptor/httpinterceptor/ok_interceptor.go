package httpinterceptor

import "context"

func OkInterceptor(_ context.Context, data any) any {
	return map[string]any{"code": 200, "msg": "成功", "data": data}
}
