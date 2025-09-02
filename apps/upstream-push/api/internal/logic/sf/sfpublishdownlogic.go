package sf

import (
	"context"

	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SfPublishDownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发布订单
func NewSfPublishDownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfPublishDownLogic {
	return &SfPublishDownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfPublishDownLogic) SfPublishDown(req *types.SfPublishDownRequest) (resp *types.EmptyType, err error) {
	// todo: add your logic here and delete this line

	return
}
