package chd

import (
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"context"

	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChdAddOnlineFeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 加小费
func NewChdAddOnlineFeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChdAddOnlineFeeLogic {
	return &ChdAddOnlineFeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChdAddOnlineFeeLogic) ChdAddOnlineFee(req *types.ChdAddOnlineFeeRequest) (resp *types.EmptyType, err error) {
	// todo: add your logic here and delete this line

	return
}
