package uu

import (
	"context"

	"chuandao-sails-core/app/api/callback/internal/svc"
	"chuandao-sails-core/app/api/callback/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UuCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单状态回调，用户提交订单时传入
func NewUuCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UuCallbackLogic {
	return &UuCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UuCallbackLogic) UuCallback(req *types.UUCallbackReq) (resp *types.UUCallbackResp, err error) {
	// todo: add your logic here and delete this line

	return
}
