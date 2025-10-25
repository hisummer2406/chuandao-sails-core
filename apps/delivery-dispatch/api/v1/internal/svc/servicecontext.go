package svc

import (
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/config"
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/middleware"
	"chuandao-sails-core/apps/platform-gateway/rpc/platformclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	HTTPLogMiddleware rest.Middleware
	SignMiddleware    rest.Middleware
	PlatformRpc       zrpc.RpcClientConf
}

func NewServiceContext(c config.Config) *ServiceContext {
	//连接platform-gateway rpc
	platformRpc := platformclient.NewPlatform(zrpc.MustNewClient(c.PlatformRpc))

	return &ServiceContext{
		Config:            c,
		HTTPLogMiddleware: middleware.NewHTTPLogMiddleware().Handle,
		SignMiddleware:    middleware.NewSignMiddleware(platformRpc).Handle,
	}
}
