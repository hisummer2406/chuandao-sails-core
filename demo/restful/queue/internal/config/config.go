package config

import (
	"chuandao-sails-core/common/rocketmq"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	RMQConfig rocketmq.Config
}
