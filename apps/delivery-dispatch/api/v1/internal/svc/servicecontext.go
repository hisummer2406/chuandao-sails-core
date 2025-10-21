package svc

import (
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/config"
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Sign   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Sign:   middleware.NewSignMiddleware().Handle,
	}
}
