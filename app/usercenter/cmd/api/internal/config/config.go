package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserCenterRpcConf zrpc.RpcClientConf
	JWT               struct {
		AccessSecret  string
		AccessExpired int64
	}
}
