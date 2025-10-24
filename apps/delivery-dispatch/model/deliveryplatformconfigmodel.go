package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DeliveryPlatformConfigModel = (*customDeliveryPlatformConfigModel)(nil)

type (
	// DeliveryPlatformConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeliveryPlatformConfigModel.
	DeliveryPlatformConfigModel interface {
		deliveryPlatformConfigModel
	}

	customDeliveryPlatformConfigModel struct {
		*defaultDeliveryPlatformConfigModel
	}
)

// NewDeliveryPlatformConfigModel returns a model for the database table.
func NewDeliveryPlatformConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DeliveryPlatformConfigModel {
	return &customDeliveryPlatformConfigModel{
		defaultDeliveryPlatformConfigModel: newDeliveryPlatformConfigModel(conn, c, opts...),
	}
}
