package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DeliveryCategoryMappingModel = (*customDeliveryCategoryMappingModel)(nil)

type (
	// DeliveryCategoryMappingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeliveryCategoryMappingModel.
	DeliveryCategoryMappingModel interface {
		deliveryCategoryMappingModel
	}

	customDeliveryCategoryMappingModel struct {
		*defaultDeliveryCategoryMappingModel
	}
)

// NewDeliveryCategoryMappingModel returns a model for the database table.
func NewDeliveryCategoryMappingModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DeliveryCategoryMappingModel {
	return &customDeliveryCategoryMappingModel{
		defaultDeliveryCategoryMappingModel: newDeliveryCategoryMappingModel(conn, c, opts...),
	}
}
