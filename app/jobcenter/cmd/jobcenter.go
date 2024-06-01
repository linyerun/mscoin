package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"mscoin/app/jobcenter/cmd/internal/config"
	"mscoin/app/jobcenter/cmd/internal/server"
	"mscoin/app/jobcenter/cmd/internal/svc"
	"os"
	"os/signal"
	"syscall"
)

var configFile = flag.String("f", "app/jobcenter/cmd/etc/jobcenter.yaml", "the config file")

func main() {
	flag.Parse()

	//日志的打印格式替换一下
	logx.MustSetup(logx.LogConf{Encoding: "plain", Stat: false})

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	jobCenterServer := server.NewJobCenterServer(ctx)

	// 优雅退出
	go func() {
		exitChan := make(chan os.Signal)
		signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-exitChan:
			logx.Info("JobCenter has been interrupted. Begin clear source!")
			jobCenterServer.Stop()
			if ctx.MongoClient != nil {
				ctx.MongoClient.Disconnect(context.Background())
			}
		}
	}()

	// 启动任务中心的定时任务
	logx.Info("Starting Job Center Server!")
	jobCenterServer.StartAndBlocking()
}
