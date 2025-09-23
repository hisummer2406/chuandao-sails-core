package chd

import (
	"context"

	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"chuandao-sails-core/apps/platform-gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChdGetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询订单
func NewChdGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChdGetOrderLogic {
	return &ChdGetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChdGetOrderLogic) ChdGetOrder(req *types.ChdGetOrderRequest) (resp *types.ChdGetOrderResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
