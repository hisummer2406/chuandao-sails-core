package fn

import (
	"context"

	"chuandao-sails-core/app/api/callback/internal/svc"
	"chuandao-sails-core/app/api/callback/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DdCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单状态回调，新增订单接口callback传入的URL
func NewDdCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DdCallbackLogic {
	return &DdCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DdCallbackLogic) DdCallback(req *types.FNCallbackReq) (resp *types.FNCallbackResp, err error) {
	// todo: add your logic here and delete this line

	return
}
