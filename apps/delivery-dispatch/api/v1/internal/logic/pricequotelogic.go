package logic

import (
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PriceQuoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询运费
func NewPriceQuoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PriceQuoteLogic {
	return &PriceQuoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PriceQuoteLogic) PriceQuote(req *types.PriceQuotaReq) (resp *types.PriceQuotaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
