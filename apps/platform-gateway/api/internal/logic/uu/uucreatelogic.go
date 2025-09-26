package uu

import (
	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"chuandao-sails-core/apps/platform-gateway/pkg/constants"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/adapter"
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
	//数据格式标准化

	return config, nil
}
