package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"mscoin/common/interceptor/grpcinterceptor"

	"mscoin/app/usercenter/cmd/rpc/internal/config"
	"mscoin/app/usercenter/cmd/rpc/internal/server"
	"mscoin/app/usercenter/cmd/rpc/internal/svc"
	"mscoin/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "app/usercenter/cmd/rpc/etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	// 防止打印过多的日志
	logx.MustSetup(logx.LogConf{Encoding: "plain", Stat: false})

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServer(grpcServer, server.NewUsercenterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 配置grpc拦截器
	s.AddUnaryInterceptors(grpcinterceptor.GlobalInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
