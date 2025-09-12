package sf

import (
	"context"

	"chuandao-sails-core/apps/upstream-gateway/api/internal/svc"
	"chuandao-sails-core/apps/upstream-gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SfAddOnlineFeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 加小费
func NewSfAddOnlineFeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfAddOnlineFeeLogic {
	return &SfAddOnlineFeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfAddOnlineFeeLogic) SfAddOnlineFee(req *types.SfAddOnlineFeeRequest) (resp *types.EmptyType, err error) {
	// todo: add your logic here and delete this line

	return
}
