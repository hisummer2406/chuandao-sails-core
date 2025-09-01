package svc

import (
	"chuandao-sails-core/common/rocketmq"
	"chuandao-sails-core/demo/restful/queue/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	RocketMQ *rocketmq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mpClient, err := rocketmq.NewClient(&c.RMQConfig)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:   c,
		RocketMQ: mpClient,
	}
}

func (ctx *ServiceContext) Stop() {
	if ctx.RocketMQ != nil {
		ctx.RocketMQ.Close()
	}
}
