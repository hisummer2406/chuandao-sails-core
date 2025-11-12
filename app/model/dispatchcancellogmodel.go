package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DispatchCancelLogModel = (*customDispatchCancelLogModel)(nil)

type (
	// DispatchCancelLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDispatchCancelLogModel.
	DispatchCancelLogModel interface {
		dispatchCancelLogModel
	}

	customDispatchCancelLogModel struct {
		*defaultDispatchCancelLogModel
	}
)

// NewDispatchCancelLogModel returns a model for the database table.
func NewDispatchCancelLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DispatchCancelLogModel {
	return &customDispatchCancelLogModel{
		defaultDispatchCancelLogModel: newDispatchCancelLogModel(conn, c, opts...),
	}
}
