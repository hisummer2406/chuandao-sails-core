package sf

import (
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"context"

	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SfGetShipperLocationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查看骑手位置
func NewSfGetShipperLocationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfGetShipperLocationLogic {
	return &SfGetShipperLocationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfGetShipperLocationLogic) SfGetShipperLocation(req *types.SfGetShipperLocationRequest) (resp *types.SfGetShipperLocationResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
