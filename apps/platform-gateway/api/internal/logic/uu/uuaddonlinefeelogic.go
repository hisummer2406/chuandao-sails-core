package uu

import (
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"context"

	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UuAddOnlineFeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 加小费
func NewUuAddOnlineFeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UuAddOnlineFeeLogic {
	return &UuAddOnlineFeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UuAddOnlineFeeLogic) UuAddOnlineFee(req *types.UUAddOnlineFeeRequest) (resp *types.EmptyType, err error) {
	// todo: add your logic here and delete this line

	return
}
