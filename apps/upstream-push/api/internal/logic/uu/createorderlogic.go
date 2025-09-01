package uu

import (
	"context"

	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 接收订单
func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderRequest) (resp *types.UUOrderResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
