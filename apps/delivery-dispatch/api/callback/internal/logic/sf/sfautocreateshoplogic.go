package sf

import (
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SfAutoCreateShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 自动建店回调
func NewSfAutoCreateShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfAutoCreateShopLogic {
	return &SfAutoCreateShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfAutoCreateShopLogic) SfAutoCreateShop() error {
	// todo: add your logic here and delete this line

	return nil
}
