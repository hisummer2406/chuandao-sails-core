package logic

import (
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCancelFeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询取消订单费用
func NewGetCancelFeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCancelFeeLogic {
	return &GetCancelFeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCancelFeeLogic) GetCancelFee(req *types.CancelOrderReq) (resp *types.GetCancelFeeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
