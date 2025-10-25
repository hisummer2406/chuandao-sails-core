package svc

import (
	"chuandao-sails-core/apps/platform-gateway/model"
	"chuandao-sails-core/apps/platform-gateway/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	PlatformModel model.PlatformConfigModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//mysql conn
	conn := sqlx.NewMysql(c.DataSource)
	//Model
	platform := model.NewPlatformConfigModel(conn, c.Redis)

	return &ServiceContext{
		Config:        c,
		PlatformModel: platform,
	}
}
