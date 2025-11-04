package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DispatchInquiryLogModel = (*customDispatchInquiryLogModel)(nil)

type (
	// DispatchInquiryLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDispatchInquiryLogModel.
	DispatchInquiryLogModel interface {
		dispatchInquiryLogModel
		InsertWithSession(ctx context.Context, session sqlx.Session, data *DispatchOrderStatusLog) (sql.Result, error)
	}

	customDispatchInquiryLogModel struct {
		*defaultDispatchInquiryLogModel
	}
)

// NewDispatchInquiryLogModel returns a model for the database table.
func NewDispatchInquiryLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DispatchInquiryLogModel {
	return &customDispatchInquiryLogModel{
		defaultDispatchInquiryLogModel: newDispatchInquiryLogModel(conn, c, opts...),
	}
}

// InsertWithSession 在事务中插入状态日志
func (m *customDispatchInquiryLogModel) InsertWithSession(ctx context.Context, session sqlx.Session, data *DispatchOrderStatusLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)",
		m.table, dispatchOrderStatusLogRowsExpectAutoSet)

	return session.ExecCtx(ctx, query,
		data.OrderNo,
		data.DeliveryCode,
		data.DeliveryOrderNo,
		data.OldStatus,
		data.NewStatus,
		data.StatusDesc,
		data.Remark,
	)
}
