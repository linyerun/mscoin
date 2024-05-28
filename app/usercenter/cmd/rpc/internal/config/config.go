package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	Captcha struct {
		Vid string
		Key string
	}
	JWT struct {
		AccessSecret  string
		AccessExpired int64
	}
	CacheRedis cache.CacheConf
}
