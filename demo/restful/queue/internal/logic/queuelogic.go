package logic

import (
	"context"

	"chuandao-sails-core/demo/restful/queue/internal/svc"
	"chuandao-sails-core/demo/restful/queue/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueueLogic {
	return &QueueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueueLogic) Queue(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
