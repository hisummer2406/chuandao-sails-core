package config

import (
	"chuandao-sails-core/common/sls"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	SLS         sls.Config
	DataSource  string
	Redis       cache.CacheConf
	PlatformRpc zrpc.RpcClientConf
}
