package logic

import (
	"context"

	"chuandao-sails-core/apps/platform-gateway/rpc/internal/svc"
	"chuandao-sails-core/apps/platform-gateway/rpc/platform"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPlatformConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPlatformConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPlatformConfigLogic {
	return &GetPlatformConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取平台配置
func (l *GetPlatformConfigLogic) GetPlatformConfig(in *platform.GetPlatformConfigReq) (*platform.GetPlatformConfigResp, error) {
	config, err := l.svcCtx.PlatformModel.FindOneByAppId(l.ctx, in.AppId)
	if err != nil {
		return nil, err
	}

	return &platform.GetPlatformConfigResp{
		AppId:        config.AppId,
		AppSecret:    config.AppSecret,
		Status:       int32(config.Status),
		PlatformName: config.AppName,
	}, nil
}
