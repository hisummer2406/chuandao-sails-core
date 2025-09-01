package v1logic

import (
	"context"

	"chuandao-sails-core/apps/delivery-strategy/rpc/internal/svc"
	"chuandao-sails-core/apps/delivery-strategy/rpc/strategy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCityConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCityConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCityConfigLogic {
	return &GetCityConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新配送策略配置
func (l *GetCityConfigLogic) GetCityConfig(in *strategy.GetCityConfigReq) (*strategy.GetCityConfigResp, error) {
	// todo: add your logic here and delete this line

	return &strategy.GetCityConfigResp{}, nil
}
