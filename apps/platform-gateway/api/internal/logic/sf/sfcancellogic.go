package sf

import (
	"context"

	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"chuandao-sails-core/apps/platform-gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SfCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单取消
func NewSfCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfCancelLogic {
	return &SfCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfCancelLogic) SfCancel(req *types.SfCancelRequest) (resp *types.SfCancelResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
