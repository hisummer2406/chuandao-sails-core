package svc

import (
	"chuandao-sails-core/apps/platform-gateway/api/internal/config"
	"chuandao-sails-core/apps/platform-gateway/api/internal/middleware"
	"chuandao-sails-core/apps/platform-gateway/model"
	"chuandao-sails-core/common/rocketmq"
	"chuandao-sails-core/common/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config            config.Config
	HTTPLogMiddleware rest.Middleware
	SignMiddleware    rest.Middleware
	PlatformModel     model.PlatformConfigModel
	MQClient          *rocketmq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	//mysql
	conn := sqlx.NewMysql(c.DataSource)

	//ID生成器
	if err := snowflake.InitDefaultGenerator(c.Snowflake.WorkerId); err != nil {
		logx.Severef("upstream-api snowflake.InitDefaultGenerator fail: %v", err)
	}

	//Model
	platform := model.NewPlatformConfigModel(conn, c.Redis)

	//MQ
	mqClient, err := rocketmq.NewClient(&c.RocketMQ)
	if err != nil {
		logx.Severef("upstream-api rocketmq connnect failed: %v", err)
	}

	logx.Infof("upstream-api rocketmq connnect ok , endpoint: %s , topics: %v", c.RocketMQ.Endpoint, c.RocketMQ.Producer.Topics)

	return &ServiceContext{
		Config:            c,
		HTTPLogMiddleware: middleware.NewHTTPLogMiddleware().Handle,
		SignMiddleware:    middleware.NewSignMiddleware(platform).Handle,
		PlatformModel:     platform,
		MQClient:          mqClient,
	}
}
