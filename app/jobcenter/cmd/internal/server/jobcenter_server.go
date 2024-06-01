package server

import (
	"context"
	"github.com/go-co-op/gocron"
	"mscoin/app/jobcenter/cmd/internal/logic"
	"mscoin/app/jobcenter/cmd/internal/svc"
	"time"
)

type JobCenterServer struct {
	ctx       *svc.ServiceContext
	scheduler *gocron.Scheduler
}

func NewJobCenterServer(ctx *svc.ServiceContext) *JobCenterServer {
	return &JobCenterServer{
		ctx:       ctx,
		scheduler: gocron.NewScheduler(time.UTC), // UTC: 和北京时间差8小时
	}
}

func (jc *JobCenterServer) run() {
	jc.scheduler.Every(1).Minute().Do(func() { logic.NewKLineLogic(context.Background(), jc.ctx).Do("1m") })
}

func (jc *JobCenterServer) StartAndBlocking() {
	jc.run()
	jc.scheduler.StartBlocking()
}

func (jc *JobCenterServer) Stop() {
	jc.scheduler.Stop()
}
