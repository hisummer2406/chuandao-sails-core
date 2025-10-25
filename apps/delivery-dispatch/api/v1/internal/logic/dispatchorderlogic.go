package logic

import (
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DispatchOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询运费后下单
func NewDispatchOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DispatchOrderLogic {
	return &DispatchOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DispatchOrderLogic) DispatchOrder(req *types.DispatchOrderReq) (resp *types.DispatchOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
