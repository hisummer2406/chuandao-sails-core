package svc

import (
	"chuandao-sails-core/apps/upstream-gateway/api/internal/config"
	"chuandao-sails-core/apps/upstream-gateway/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config            config.Config
	HTTPLogMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		HTTPLogMiddleware: middleware.NewHTTPLogMiddleware().Handle,
	}
}
