package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"mscoin/app/usercenter/cmd/rpc/internal/config"
	"mscoin/common/dbinit"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// redis cache client
	redisCache := cache.New(c.CacheRedis, nil, cache.NewStat("mscoin"), nil, func(o *cache.Options) {})

	// mysql client
	gormMysqlClient, err := dbinit.CreateGormMysqlClient(c.Mysql.DataSource, 100, 10)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
		Db:     gormMysqlClient,
	}
}
