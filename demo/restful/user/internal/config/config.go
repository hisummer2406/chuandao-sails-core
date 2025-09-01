package config

import (
	"chuandao-sails-core/common/rocketmq"
	"chuandao-sails-core/common/sls"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	SLS      sls.Config
	RocketMQ rocketmq.Config
}
