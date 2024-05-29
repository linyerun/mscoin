package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"mscoin/app/usercenter/cmd/api/internal/config"
	"mscoin/app/usercenter/cmd/api/internal/handler"
	"mscoin/app/usercenter/cmd/api/internal/svc"
	"mscoin/common/interceptor/httpinterceptor"
)

var configFile = flag.String("f", "app/usercenter/cmd/api/etc/user_center_api.yaml", "the config file")

func main() {
	flag.Parse()

	// 防止打印过多的日志
	logx.MustSetup(logx.LogConf{Encoding: "plain", Stat: false})

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 配置全局成功处理
	httpx.SetOkHandler(httpinterceptor.OkInterceptor)

	// 配置全局异常处理
	httpx.SetErrorHandler(httpinterceptor.ErrorInterceptor)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
