package uu

import (
	"context"

	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"chuandao-sails-core/apps/platform-gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单详情
func NewUuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UuDetailLogic {
	return &UuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UuDetailLogic) UuDetail(req *types.UUOrderDetailRequest) (resp *types.UUOrderDetailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
