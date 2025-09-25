package uu

import (
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"context"

	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UuCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 取消订单
func NewUuCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UuCancelLogic {
	return &UuCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UuCancelLogic) UuCancel(req *types.UUCancelOrderRequest) (resp *types.EmptyType, err error) {
	// todo: add your logic here and delete this line

	return
}
