package v1logic

import (
	"context"

	"chuandao-sails-core/apps/delivery-strategy/rpc/internal/svc"
	"chuandao-sails-core/apps/delivery-strategy/rpc/strategy"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectBestPlatformLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectBestPlatformLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectBestPlatformLogic {
	return &SelectBestPlatformLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 选择最优配送平台
func (l *SelectBestPlatformLogic) SelectBestPlatform(in *strategy.SelectBestPlatformReq) (*strategy.SelectBestPlatformResp, error) {
	// todo: add your logic here and delete this line

	return &strategy.SelectBestPlatformResp{}, nil
}
