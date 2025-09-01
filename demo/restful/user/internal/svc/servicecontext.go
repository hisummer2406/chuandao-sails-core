package svc

import (
	"chuandao-sails-core/common/rocketmq"
	"chuandao-sails-core/demo/restful/user/internal/config"
	"chuandao-sails-core/demo/restful/user/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"log"
)

type ServiceContext struct {
	Config            config.Config
	HTTPLogMiddleware rest.Middleware
	RocketMQ          *rocketmq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	//初始化rocketmq
	mpClient, err := rocketmq.NewClient(&c.RocketMQ)
	if err != nil {
		log.Fatalf("new rocketmq client failed: %v", err)
	}

	return &ServiceContext{
		Config:            c,
		HTTPLogMiddleware: middleware.HTTPLogMiddleware(),
		RocketMQ:          mpClient,
	}
}

func (svc *ServiceContext) Stop() {
	if svc.RocketMQ != nil {
		svc.RocketMQ.Close()
	}
}
