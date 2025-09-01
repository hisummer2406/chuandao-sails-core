package v1logic

import (
	"context"

	"chuandao-sails-core/apps/delivery-strategy/rpc/internal/svc"
	"chuandao-sails-core/apps/delivery-strategy/rpc/strategy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStrategyConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStrategyConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStrategyConfigLogic {
	return &GetStrategyConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取配送策略配置
func (l *GetStrategyConfigLogic) GetStrategyConfig(in *strategy.GetStrategyConfigReq) (*strategy.GetStrategyConfigResp, error) {
	// todo: add your logic here and delete this line

	return &strategy.GetStrategyConfigResp{}, nil
}
