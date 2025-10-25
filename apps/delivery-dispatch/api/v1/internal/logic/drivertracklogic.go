package logic

import (
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DriverTrackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查看跑男位置
func NewDriverTrackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DriverTrackLogic {
	return &DriverTrackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DriverTrackLogic) DriverTrack(req *types.QueryOrderReq) (resp *types.DrickerTrackResp, err error) {
	// todo: add your logic here and delete this line

	return
}
