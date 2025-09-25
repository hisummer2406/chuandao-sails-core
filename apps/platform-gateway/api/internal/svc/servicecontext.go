package svc

import (
	"chuandao-sails-core/apps/platform-gateway/api/internal/config"
	"chuandao-sails-core/apps/platform-gateway/api/internal/middleware"
	"chuandao-sails-core/apps/platform-gateway/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config            config.Config
	HTTPLogMiddleware rest.Middleware
	SignMiddleware    rest.Middleware
	PlatformModel     model.PlatformConfigModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//mysql
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:            c,
		HTTPLogMiddleware: middleware.NewHTTPLogMiddleware().Handle,
		SignMiddleware:    middleware.NewSignMiddleware().Handle,
		PlatformModel:     model.NewPlatformConfigModel(conn, c.Redis),
	}
}
