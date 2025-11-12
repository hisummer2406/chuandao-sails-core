package logic

import (
	"context"

	"chuandao-sails-core/app/api/delivery/internal/svc"
	"chuandao-sails-core/app/api/delivery/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDriverLocationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查看跑男位置
func NewGetDriverLocationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDriverLocationLogic {
	return &GetDriverLocationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDriverLocationLogic) GetDriverLocation(req *types.QueryOrderReq) (resp *types.GetDriverLocResp, err error) {
	// todo: add your logic here and delete this line

	return
}
