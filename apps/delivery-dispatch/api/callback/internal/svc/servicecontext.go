package svc

import (
	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
