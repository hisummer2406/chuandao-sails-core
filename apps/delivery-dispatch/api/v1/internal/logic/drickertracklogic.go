package logic

import (
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/types"
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DrickerTrackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查看跑男位置
func NewDrickerTrackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrickerTrackLogic {
	return &DrickerTrackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DrickerTrackLogic) DrickerTrack(req *types.QueryOrderReq) (resp *types.DrickerTrackResp, err error) {
	// todo: add your logic here and delete this line

	return
}
