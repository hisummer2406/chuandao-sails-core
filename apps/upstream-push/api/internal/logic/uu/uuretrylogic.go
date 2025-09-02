package uu

import (
	"context"

	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UuRetryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单回推
func NewUuRetryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UuRetryLogic {
	return &UuRetryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UuRetryLogic) UuRetry(req *types.UURetryRequest) (resp *types.EmptyType, err error) {
	// todo: add your logic here and delete this line

	return
}
