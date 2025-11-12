package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TrackOrderTrackLogModel = (*customTrackOrderTrackLogModel)(nil)

type (
	// TrackOrderTrackLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTrackOrderTrackLogModel.
	TrackOrderTrackLogModel interface {
		trackOrderTrackLogModel
	}

	customTrackOrderTrackLogModel struct {
		*defaultTrackOrderTrackLogModel
	}
)

// NewTrackOrderTrackLogModel returns a model for the database table.
func NewTrackOrderTrackLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TrackOrderTrackLogModel {
	return &customTrackOrderTrackLogModel{
		defaultTrackOrderTrackLogModel: newTrackOrderTrackLogModel(conn, c, opts...),
	}
}
