package sf

import (
	"context"

	"chuandao-sails-core/app/api/callback/internal/svc"
	"chuandao-sails-core/app/api/callback/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SfOrderStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 配送状态更改回调
func NewSfOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfOrderStatusLogic {
	return &SfOrderStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfOrderStatusLogic) SfOrderStatus(req *types.SFOrderStatusReq) (resp *types.SFResp, err error) {
	// todo: add your logic here and delete this line

	return
}
