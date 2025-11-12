package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DeliveryVirtualShopModel = (*customDeliveryVirtualShopModel)(nil)

type (
	// DeliveryVirtualShopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeliveryVirtualShopModel.
	DeliveryVirtualShopModel interface {
		deliveryVirtualShopModel
	}

	customDeliveryVirtualShopModel struct {
		*defaultDeliveryVirtualShopModel
	}
)

// NewDeliveryVirtualShopModel returns a model for the database table.
func NewDeliveryVirtualShopModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DeliveryVirtualShopModel {
	return &customDeliveryVirtualShopModel{
		defaultDeliveryVirtualShopModel: newDeliveryVirtualShopModel(conn, c, opts...),
	}
}
