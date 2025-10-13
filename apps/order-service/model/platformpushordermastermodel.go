package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PlatformPushOrderMasterModel = (*customPlatformPushOrderMasterModel)(nil)

type (
	// PlatformPushOrderMasterModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlatformPushOrderMasterModel.
	PlatformPushOrderMasterModel interface {
		platformPushOrderMasterModel
	}

	customPlatformPushOrderMasterModel struct {
		*defaultPlatformPushOrderMasterModel
	}
)

// NewPlatformPushOrderMasterModel returns a model for the database table.
func NewPlatformPushOrderMasterModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PlatformPushOrderMasterModel {
	return &customPlatformPushOrderMasterModel{
		defaultPlatformPushOrderMasterModel: newPlatformPushOrderMasterModel(conn, c, opts...),
	}
}
