package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TrackPlatformCallbackLogModel = (*customTrackPlatformCallbackLogModel)(nil)

type (
	// TrackPlatformCallbackLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTrackPlatformCallbackLogModel.
	TrackPlatformCallbackLogModel interface {
		trackPlatformCallbackLogModel
	}

	customTrackPlatformCallbackLogModel struct {
		*defaultTrackPlatformCallbackLogModel
	}
)

// NewTrackPlatformCallbackLogModel returns a model for the database table.
func NewTrackPlatformCallbackLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TrackPlatformCallbackLogModel {
	return &customTrackPlatformCallbackLogModel{
		defaultTrackPlatformCallbackLogModel: newTrackPlatformCallbackLogModel(conn, c, opts...),
	}
}
