package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DispatchOrderStatusLogModel = (*customDispatchOrderStatusLogModel)(nil)

type (
	// DispatchOrderStatusLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDispatchOrderStatusLogModel.
	DispatchOrderStatusLogModel interface {
		dispatchOrderStatusLogModel
	}

	customDispatchOrderStatusLogModel struct {
		*defaultDispatchOrderStatusLogModel
	}
)

// NewDispatchOrderStatusLogModel returns a model for the database table.
func NewDispatchOrderStatusLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DispatchOrderStatusLogModel {
	return &customDispatchOrderStatusLogModel{
		defaultDispatchOrderStatusLogModel: newDispatchOrderStatusLogModel(conn, c, opts...),
	}
}
