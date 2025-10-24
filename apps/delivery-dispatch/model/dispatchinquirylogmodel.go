package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DispatchInquiryLogModel = (*customDispatchInquiryLogModel)(nil)

type (
	// DispatchInquiryLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDispatchInquiryLogModel.
	DispatchInquiryLogModel interface {
		dispatchInquiryLogModel
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
