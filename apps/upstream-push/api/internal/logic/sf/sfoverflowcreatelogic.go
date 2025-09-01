package sf

import (
	"context"

	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SfOverflowCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单创建
func NewSfOverflowCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfOverflowCreateLogic {
	return &SfOverflowCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfOverflowCreateLogic) SfOverflowCreate(req *types.SfOverflowCreateRequest) (resp *types.SfOverflowCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
