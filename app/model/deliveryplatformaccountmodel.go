package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DeliveryPlatformAccountModel = (*customDeliveryPlatformAccountModel)(nil)

type (
	// DeliveryPlatformAccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeliveryPlatformAccountModel.
	DeliveryPlatformAccountModel interface {
		deliveryPlatformAccountModel
	}

	customDeliveryPlatformAccountModel struct {
		*defaultDeliveryPlatformAccountModel
	}
)

// NewDeliveryPlatformAccountModel returns a model for the database table.
func NewDeliveryPlatformAccountModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DeliveryPlatformAccountModel {
	return &customDeliveryPlatformAccountModel{
		defaultDeliveryPlatformAccountModel: newDeliveryPlatformAccountModel(conn, c, opts...),
	}
}
