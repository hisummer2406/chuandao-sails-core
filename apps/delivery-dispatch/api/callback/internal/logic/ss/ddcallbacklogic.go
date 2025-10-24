package ss

import (
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DdCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单状态回调，回调账号配置的地址（账户中心->应用中心->回调URL：NotifyUrl）来通知订单状态的变化。
func NewDdCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DdCallbackLogic {
	return &DdCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DdCallbackLogic) DdCallback(req *types.SSCallbackReq) (resp *types.SSCallbackResp, err error) {
	// todo: add your logic here and delete this line

	return
}
