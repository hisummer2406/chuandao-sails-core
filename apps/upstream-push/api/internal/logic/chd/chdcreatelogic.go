package chd

import (
	"context"

	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChdCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单创建
func NewChdCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChdCreateLogic {
	return &ChdCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChdCreateLogic) ChdCreate(req *types.ChdCreateRequest) (resp *types.ChdCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
