package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DispatchInquiryDetailModel = (*customDispatchInquiryDetailModel)(nil)

type (
	// DispatchInquiryDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDispatchInquiryDetailModel.
	DispatchInquiryDetailModel interface {
		dispatchInquiryDetailModel
	}

	customDispatchInquiryDetailModel struct {
		*defaultDispatchInquiryDetailModel
	}
)

// NewDispatchInquiryDetailModel returns a model for the database table.
func NewDispatchInquiryDetailModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DispatchInquiryDetailModel {
	return &customDispatchInquiryDetailModel{
		defaultDispatchInquiryDetailModel: newDispatchInquiryDetailModel(conn, c, opts...),
	}
}
