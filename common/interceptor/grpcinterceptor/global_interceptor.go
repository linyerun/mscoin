package grpcinterceptor

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mscoin/common/xerr"
)

// GlobalInterceptor 说明
// 1. 全程不在logic打日志，全部在这个拦截器里面把日志输出
// 2. 把error从xerr到grpc的error格式进行传递
func GlobalInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 打印请求参数
	logx.Infof("req args: %+v", req)

	// 执行请求
	resp, err = handler(ctx, req)
	if err != nil {
		// 打印错误信息
		logx.Error(err)

		// xerr => grpc err
		if _err, ok := errors.Cause(err).(xerr.Error); ok { //被errors包装了一层
			return nil, status.Error(codes.Code(_err.GetCode()), _err.GetMsg())
		} else if _err, ok = err.(xerr.Error); ok { // 没被包装
			return nil, status.Error(codes.Code(_err.GetCode()), _err.GetMsg())
		}

		// 未知err
		return nil, status.Error(codes.Code(xerr.ServerError), err.Error())
	}

	return resp, nil
}
