package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DispatchOrderModel = (*customDispatchOrderModel)(nil)

type (
	// DispatchOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDispatchOrderModel.
	DispatchOrderModel interface {
		dispatchOrderModel
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		InsertWithSessions(ctx context.Context, session sqlx.Session, data *DispatchOrder) (sql.Result, error)
	}

	customDispatchOrderModel struct {
		*defaultDispatchOrderModel
	}
)

// NewDispatchOrderModel returns a model for the database table.
func NewDispatchOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DispatchOrderModel {
	return &customDispatchOrderModel{
		defaultDispatchOrderModel: newDispatchOrderModel(conn, c, opts...),
	}
}

// Trans 事务封装
func (m *customDispatchOrderModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, fn)
}

// InsertWithSessions 在事务中插入
func (m *customDispatchOrderModel) InsertWithSessions(ctx context.Context, session sqlx.Session, data *DispatchOrder) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		m.table, dispatchOrderRowsExpectAutoSet)

	return session.ExecCtx(ctx, query,
		data.OrderNo,
		data.OriginOrderId,
		data.UpstreamSource,
		data.UpstreamOrderId,
		data.DeliveryCode,
		data.DeliveryOrderNo,
		data.ShortNum,
		data.AccountId,
		data.SourceAppId,
		data.InquiryId,
		data.Status,
		data.FromMobile,
		data.FromAddress,
		data.ToMobile,
		data.ToAddress,
		data.Note,
		data.GoodsName,
		data.GoodsType,
		data.GoodsDetail,
		data.TotalAmount,
		data.PriceDetail,
		data.Distance,
		data.DriverName,
		data.DriverMobile,
		data.ShopId,
		data.SubscribeType,
		data.SubscribeTime,
		data.DisableDelivery,
	)
}
