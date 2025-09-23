package uu

import (
	"context"

	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"chuandao-sails-core/apps/platform-gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UuGatewayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 接收订单
func NewUuGatewayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UuGatewayLogic {
	return &UuGatewayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UuGatewayLogic) UuGateway(req *types.UUCreateOrderRequest) (resp *types.EmptyType, err error) {
	// todo: add your logic here and delete this line

	return
}
