package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DispatchFeeAddLogModel = (*customDispatchFeeAddLogModel)(nil)

type (
	// DispatchFeeAddLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDispatchFeeAddLogModel.
	DispatchFeeAddLogModel interface {
		dispatchFeeAddLogModel
	}

	customDispatchFeeAddLogModel struct {
		*defaultDispatchFeeAddLogModel
	}
)

// NewDispatchFeeAddLogModel returns a model for the database table.
func NewDispatchFeeAddLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DispatchFeeAddLogModel {
	return &customDispatchFeeAddLogModel{
		defaultDispatchFeeAddLogModel: newDispatchFeeAddLogModel(conn, c, opts...),
	}
}
