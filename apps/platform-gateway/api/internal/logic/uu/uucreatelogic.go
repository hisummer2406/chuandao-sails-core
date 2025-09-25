package uu

import (
	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UuCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUuCreateLogic 接收订单
func NewUuCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UuCreateLogic {
	return &UuCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UuCreateLogic) UuCreate(req *types.UUCreateOrderRequest) (resp interface{}, err error) {
	fmt.Printf("UUCreateOrderRequest appId: %s , %s", req.AppId, req.Sign)
	config, err := l.svcCtx.PlatformModel.FindOneByAppId(l.ctx, req.AppId)

	if err != nil {
		logc.Error(l.ctx, "UuCreateLogic FindOneByAppId error", err)
		return nil, err
	}

	return config, nil
}
