package sf

import (
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"context"

	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SfGetPriceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 询价
func NewSfGetPriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfGetPriceLogic {
	return &SfGetPriceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfGetPriceLogic) SfGetPrice(req *types.SfGetPriceRequest) (resp *types.SfGetPriceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
