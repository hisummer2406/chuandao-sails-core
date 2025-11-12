package sf

import (
	"context"

	"chuandao-sails-core/app/api/callback/internal/svc"
	"chuandao-sails-core/app/api/callback/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SfOrderExceptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单配送异常
func NewSfOrderExceptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfOrderExceptionLogic {
	return &SfOrderExceptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfOrderExceptionLogic) SfOrderException(req *types.SFOrderExceptionReq) error {
	// todo: add your logic here and delete this line

	return nil
}
