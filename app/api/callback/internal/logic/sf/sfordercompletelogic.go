package sf

import (
	"context"

	"chuandao-sails-core/app/api/callback/internal/svc"
	"chuandao-sails-core/app/api/callback/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SfOrderCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单完成回调
func NewSfOrderCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfOrderCompleteLogic {
	return &SfOrderCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfOrderCompleteLogic) SfOrderComplete(req *types.SFOrderCompleteReq) (resp *types.SFResp, err error) {
	// todo: add your logic here and delete this line

	return
}
