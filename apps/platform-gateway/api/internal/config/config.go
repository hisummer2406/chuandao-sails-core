package config

import (
	"chuandao-sails-core/common/sls"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DataSource string
	SLS        sls.Config
	Redis      cache.CacheConf
}
