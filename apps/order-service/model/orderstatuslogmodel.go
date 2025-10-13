package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderStatusLogModel = (*customOrderStatusLogModel)(nil)

type (
	// OrderStatusLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderStatusLogModel.
	OrderStatusLogModel interface {
		orderStatusLogModel
	}

	customOrderStatusLogModel struct {
		*defaultOrderStatusLogModel
	}
)

// NewOrderStatusLogModel returns a model for the database table.
func NewOrderStatusLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderStatusLogModel {
	return &customOrderStatusLogModel{
		defaultOrderStatusLogModel: newOrderStatusLogModel(conn, c, opts...),
	}
}
