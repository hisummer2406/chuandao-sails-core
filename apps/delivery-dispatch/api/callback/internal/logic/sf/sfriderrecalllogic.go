package sf

import (
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SfRiderRecallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 骑士撤单状态回调
func NewSfRiderRecallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfRiderRecallLogic {
	return &SfRiderRecallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfRiderRecallLogic) SfRiderRecall(req *types.SFRiderRecallReq) (resp *types.SFResp, err error) {
	// todo: add your logic here and delete this line

	return
}
