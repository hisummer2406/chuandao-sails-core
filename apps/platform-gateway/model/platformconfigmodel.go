package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PlatformConfigModel = (*customPlatformConfigModel)(nil)

type (
	// PlatformConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlatformConfigModel.
	PlatformConfigModel interface {
		platformConfigModel
		FindOneByAppId(ctx context.Context, appId string) (*PlatformConfig, error)
	}

	customPlatformConfigModel struct {
		*defaultPlatformConfigModel
	}
)

// NewPlatformConfigModel returns a model for the database table.
func NewPlatformConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PlatformConfigModel {
	return &customPlatformConfigModel{
		defaultPlatformConfigModel: newPlatformConfigModel(conn, c, opts...),
	}
}

func (m *customPlatformConfigModel) FindOneByAppId(ctx context.Context, appId string) (*PlatformConfig, error) {
	cacheCdtsPlatformDbPlatformConfigAppIdKey := fmt.Sprintf("cache:cdtsPlatformDb:platformConfig:appId:%s", appId)
	var resp PlatformConfig
	err := m.QueryRowIndexCtx(ctx, &resp, cacheCdtsPlatformDbPlatformConfigAppIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `app_id` = ? limit 1", platformConfigRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, appId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
