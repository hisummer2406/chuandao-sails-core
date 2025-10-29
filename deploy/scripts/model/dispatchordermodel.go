package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DispatchOrderModel = (*customDispatchOrderModel)(nil)

type (
	// DispatchOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDispatchOrderModel.
	DispatchOrderModel interface {
		dispatchOrderModel
	}

	customDispatchOrderModel struct {
		*defaultDispatchOrderModel
	}
)

// NewDispatchOrderModel returns a model for the database table.
func NewDispatchOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DispatchOrderModel {
	return &customDispatchOrderModel{
		defaultDispatchOrderModel: newDispatchOrderModel(conn, c, opts...),
	}
}
