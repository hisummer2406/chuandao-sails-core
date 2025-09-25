package config

import (
	"chuandao-sails-core/common/sls"
	"chuandao-sails-core/common/snowflake"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DataSource string
	SLS        sls.Config
	Redis      cache.CacheConf
	Snowflake  snowflake.Config
}
