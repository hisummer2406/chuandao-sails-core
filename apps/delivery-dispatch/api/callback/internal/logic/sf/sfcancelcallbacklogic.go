package sf

import (
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SfCancelCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 顺丰原因订单取消回调
func NewSfCancelCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfCancelCallbackLogic {
	return &SfCancelCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfCancelCallbackLogic) SfCancelCallback(req *types.SFCancelCallbackReq) (resp *types.SFResp, err error) {
	// todo: add your logic here and delete this line

	return
}
