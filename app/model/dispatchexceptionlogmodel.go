package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DispatchExceptionLogModel = (*customDispatchExceptionLogModel)(nil)

type (
	// DispatchExceptionLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDispatchExceptionLogModel.
	DispatchExceptionLogModel interface {
		dispatchExceptionLogModel
	}

	customDispatchExceptionLogModel struct {
		*defaultDispatchExceptionLogModel
	}
)

// NewDispatchExceptionLogModel returns a model for the database table.
func NewDispatchExceptionLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DispatchExceptionLogModel {
	return &customDispatchExceptionLogModel{
		defaultDispatchExceptionLogModel: newDispatchExceptionLogModel(conn, c, opts...),
	}
}
