package logic

import (
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/types"
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPenaltyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询取消订单费用
func NewQueryPenaltyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPenaltyLogic {
	return &QueryPenaltyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPenaltyLogic) QueryPenalty(req *types.CancelOrderReq) (resp *types.QueryPenaltyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
