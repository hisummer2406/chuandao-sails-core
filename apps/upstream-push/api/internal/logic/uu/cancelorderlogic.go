package uu

import (
	"context"

	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 接收取消订单消息
func NewCancelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOrderLogic {
	return &CancelOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelOrderLogic) CancelOrder(req *types.CancelOrderRequest) (resp *types.UUOrderResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
