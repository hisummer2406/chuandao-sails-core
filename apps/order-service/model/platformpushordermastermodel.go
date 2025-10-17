package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 编译时接口检查 的常见写法，确保类型实现了指定的接口。
var _ PlatformPushOrderMasterModel = (*customPlatformPushOrderMasterModel)(nil)

type (
	// PlatformPushOrderMasterModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlatformPushOrderMasterModel.
	PlatformPushOrderMasterModel interface {
		platformPushOrderMasterModel
		//自定义方法
		FindOneByOrderId(ctx context.Context, orderId string) (*PlatformPushOrderMaster, error)
		UpdateStatus(ctx context.Context, orderId string, oldStatus, newStatus string) error
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	}

	customPlatformPushOrderMasterModel struct {
		*defaultPlatformPushOrderMasterModel
	}
)

// NewPlatformPushOrderMasterModel returns a model for the database table.
func NewPlatformPushOrderMasterModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PlatformPushOrderMasterModel {
	return &customPlatformPushOrderMasterModel{
		defaultPlatformPushOrderMasterModel: newPlatformPushOrderMasterModel(conn, c, opts...),
	}
}

// FindOneByOrderId 订单号查询订单
func (m *customPlatformPushOrderMasterModel) FindOneByOrderId(ctx context.Context, orderId string) (*PlatformPushOrderMaster, error) {
	cacheKey := fmt.Sprintf("cache:orderId:%s", orderId)
	var resp PlatformPushOrderMaster

	err := m.QueryRowIndexCtx(ctx, &resp, cacheKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `order_id` = ? limit 1", platformPushOrderMasterRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, orderId); err != nil {
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

// UpdateStatus 更新订单状态
func (m *customPlatformPushOrderMasterModel) UpdateStatus(ctx context.Context, orderId string, oldStatus, newStatus string) error {
	//先查询获取主键ID用于缓存删除
	order, err := m.FindOneByOrderId(ctx, orderId)
	if err != nil {
		return err
	}

	cacheKey := m.formatPrimary(order.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `status` = ? where `order_id` = ?  and `status` = ?", m.table)
		return conn.ExecCtx(ctx, query, newStatus, orderId, oldStatus)
	}, cacheKey)

	return err
}

// Trans 事务支持
func (m *customPlatformPushOrderMasterModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, fn)
}
