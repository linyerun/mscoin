package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"mscoin/app/usercenter/cmd/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(c.CacheRedis, nil, cache.NewStat("mscoin"), nil, func(o *cache.Options) {})

	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
		Db:     gormMysqlInit(c.Mysql.DataSource),
	}
}
