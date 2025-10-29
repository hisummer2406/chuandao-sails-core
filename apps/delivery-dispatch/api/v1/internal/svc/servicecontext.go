package svc

import (
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/config"
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/middleware"
	"chuandao-sails-core/apps/delivery-dispatch/model"
	"chuandao-sails-core/apps/delivery-dispatch/pkg/engine/platform"
	"chuandao-sails-core/apps/delivery-dispatch/pkg/engine/pricing"
	"chuandao-sails-core/apps/platform-gateway/rpc/platformclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	HTTPLogMiddleware rest.Middleware
	SignMiddleware    rest.Middleware
	PlatformRpc       zrpc.RpcClientConf

	PlatformManager *platform.PlatformManager //平台管理器
	PricingEngine   *pricing.PricingEngine    //询价引擎

	DispatchOrderModel          model.DispatchOrderModel //订单表
	DispatchOrderStatusLogModel model.DispatchOrderStatusLogModel
	DispatchInquiryLogModel     model.DispatchInquiryLogModel //询价
	DispatchInquiryDetailModel  model.DispatchInquiryDetailModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//连接platform-gateway rpc
	platformRpc := platformclient.NewPlatform(zrpc.MustNewClient(c.PlatformRpc))

	//model
	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:                      c,
		HTTPLogMiddleware:           middleware.NewHTTPLogMiddleware().Handle,
		SignMiddleware:              middleware.NewSignMiddleware(platformRpc).Handle,
		DispatchOrderModel:          model.NewDispatchOrderModel(conn, c.Redis),
		DispatchOrderStatusLogModel: model.NewDispatchOrderStatusLogModel(conn, c.Redis),
		DispatchInquiryLogModel:     model.NewDispatchInquiryLogModel(conn, c.Redis),
		DispatchInquiryDetailModel:  model.NewDispatchInquiryDetailModel(conn, c.Redis),
	}
}
