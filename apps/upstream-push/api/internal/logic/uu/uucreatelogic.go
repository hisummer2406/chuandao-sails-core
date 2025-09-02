package uu

import (
	"context"

	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UuCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 接收订单
func NewUuCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UuCreateLogic {
	return &UuCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UuCreateLogic) UuCreate(req *types.UUCreateOrderRequest) (resp *types.EmptyType, err error) {
	// todo: add your logic here and delete this line

	return
}
