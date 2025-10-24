package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DispatchRetryLogModel = (*customDispatchRetryLogModel)(nil)

type (
	// DispatchRetryLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDispatchRetryLogModel.
	DispatchRetryLogModel interface {
		dispatchRetryLogModel
	}

	customDispatchRetryLogModel struct {
		*defaultDispatchRetryLogModel
	}
)

// NewDispatchRetryLogModel returns a model for the database table.
func NewDispatchRetryLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DispatchRetryLogModel {
	return &customDispatchRetryLogModel{
		defaultDispatchRetryLogModel: newDispatchRetryLogModel(conn, c, opts...),
	}
}
