package logic

import (
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询订单详情
func NewQueryOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderLogic {
	return &QueryOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOrderLogic) QueryOrder(req *types.QueryOrderReq) (resp *types.QueryOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
