package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StrategyDispatchConfigModel = (*customStrategyDispatchConfigModel)(nil)

type (
	// StrategyDispatchConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStrategyDispatchConfigModel.
	StrategyDispatchConfigModel interface {
		strategyDispatchConfigModel
	}

	customStrategyDispatchConfigModel struct {
		*defaultStrategyDispatchConfigModel
	}
)

// NewStrategyDispatchConfigModel returns a model for the database table.
func NewStrategyDispatchConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StrategyDispatchConfigModel {
	return &customStrategyDispatchConfigModel{
		defaultStrategyDispatchConfigModel: newStrategyDispatchConfigModel(conn, c, opts...),
	}
}
