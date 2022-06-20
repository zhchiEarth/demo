package server

import (
	"compound/internal/conf"
	"compound/internal/task"
	"context"
	"github.com/go-co-op/gocron"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type CronServer struct {
	s   *gocron.Scheduler
	err error
	log *log.Helper
	//tasks []task.Job
}

func NewCronServer(c *conf.Server, logger log.Logger, market *task.MarketTask, account *task.AccountTask,
	oracle *task.OraclePriceTask) *CronServer {
	srv := &CronServer{
		log: log.NewHelper(log.With(logger, "module", "server/cron")),
	}
	l, _ := time.LoadLocation("Asia/Shanghai")
	srv.s = gocron.NewScheduler(l)

	// 同步市场 1分钟一次
	job, err := srv.s.Every(1).Minute().Do(func() {
		market.Handle()
	})
	if err != nil {
		return nil
	}
	market.SetJob(job)

	// 用户同步
	job, err = srv.s.Every(30).Second().Do(func() {
		account.Handle()
	})
	if err != nil {
		return nil
	}
	account.SetJob(job)

	// 同步预言机价格
	job, err = srv.s.Every(30).Second().Do(func() {
		oracle.Handle()
	})
	if err != nil {
		return nil
	}
	oracle.SetJob(job)

	srv.s.StartAsync()
	return srv
}

// Start the Discord server.
func (cron *CronServer) Start(ctx context.Context) error {
	if cron.err != nil {
		return cron.err
	}
	cron.log.Infof("[Cron] server listening on:")
	return nil
}

// Stop the ws server.
func (cron *CronServer) Stop(ctx context.Context) error {
	cron.log.Info("[Cron] server stopping")
	cron.s.Stop()
	return nil
}
