package chd

import (
	"context"

	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OverflowCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单创建
func NewOverflowCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OverflowCreateLogic {
	return &OverflowCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OverflowCreateLogic) OverflowCreate(req *types.OverflowCreateRequest) (resp *types.OverflowCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
