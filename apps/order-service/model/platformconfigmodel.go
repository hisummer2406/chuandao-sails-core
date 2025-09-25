package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PlatformConfigModel = (*customPlatformConfigModel)(nil)

type (
	// PlatformConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlatformConfigModel.
	PlatformConfigModel interface {
		platformConfigModel
	}

	customPlatformConfigModel struct {
		*defaultPlatformConfigModel
	}
)

// NewPlatformConfigModel returns a model for the database table.
func NewPlatformConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PlatformConfigModel {
	return &customPlatformConfigModel{
		defaultPlatformConfigModel: newPlatformConfigModel(conn, c, opts...),
	}
}
