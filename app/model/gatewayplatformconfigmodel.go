package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GatewayPlatformConfigModel = (*customGatewayPlatformConfigModel)(nil)

type (
	// GatewayPlatformConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGatewayPlatformConfigModel.
	GatewayPlatformConfigModel interface {
		gatewayPlatformConfigModel
	}

	customGatewayPlatformConfigModel struct {
		*defaultGatewayPlatformConfigModel
	}
)

// NewGatewayPlatformConfigModel returns a model for the database table.
func NewGatewayPlatformConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GatewayPlatformConfigModel {
	return &customGatewayPlatformConfigModel{
		defaultGatewayPlatformConfigModel: newGatewayPlatformConfigModel(conn, c, opts...),
	}
}
