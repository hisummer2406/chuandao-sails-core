package logic

import (
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/delivery/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/delivery/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 加小费
func NewAddTipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTipLogic {
	return &AddTipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTipLogic) AddTip(req *types.AddTipReq) (resp *types.AddTipResp, err error) {
	// todo: add your logic here and delete this line

	return
}
