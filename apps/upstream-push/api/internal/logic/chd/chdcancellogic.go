package chd

import (
	"context"

	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChdCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单取消
func NewChdCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChdCancelLogic {
	return &ChdCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChdCancelLogic) ChdCancel(req *types.ChdCancelRequest) (resp *types.EmptyType, err error) {
	// todo: add your logic here and delete this line

	return
}
