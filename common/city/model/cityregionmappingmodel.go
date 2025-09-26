package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CityRegionMappingModel = (*customCityRegionMappingModel)(nil)

type (
	// CityRegionMappingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCityRegionMappingModel.
	CityRegionMappingModel interface {
		cityRegionMappingModel
	}

	customCityRegionMappingModel struct {
		*defaultCityRegionMappingModel
	}
)

// NewCityRegionMappingModel returns a model for the database table.
func NewCityRegionMappingModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CityRegionMappingModel {
	return &customCityRegionMappingModel{
		defaultCityRegionMappingModel: newCityRegionMappingModel(conn, c, opts...),
	}
}
