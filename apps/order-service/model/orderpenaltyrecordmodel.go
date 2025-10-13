package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderPenaltyRecordModel = (*customOrderPenaltyRecordModel)(nil)

type (
	// OrderPenaltyRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderPenaltyRecordModel.
	OrderPenaltyRecordModel interface {
		orderPenaltyRecordModel
	}

	customOrderPenaltyRecordModel struct {
		*defaultOrderPenaltyRecordModel
	}
)

// NewOrderPenaltyRecordModel returns a model for the database table.
func NewOrderPenaltyRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderPenaltyRecordModel {
	return &customOrderPenaltyRecordModel{
		defaultOrderPenaltyRecordModel: newOrderPenaltyRecordModel(conn, c, opts...),
	}
}
