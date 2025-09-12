package sf

import (
	"context"

	"chuandao-sails-core/apps/upstream-gateway/api/internal/svc"
	"chuandao-sails-core/apps/upstream-gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SfCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单创建
func NewSfCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfCreateLogic {
	return &SfCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfCreateLogic) SfCreate(req *types.SfCreateRequest) (resp *types.SfCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
