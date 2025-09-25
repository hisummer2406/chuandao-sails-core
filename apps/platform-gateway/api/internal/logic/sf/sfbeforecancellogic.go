package sf

import (
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"context"

	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SfBeforeCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 预取消
func NewSfBeforeCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SfBeforeCancelLogic {
	return &SfBeforeCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SfBeforeCancelLogic) SfBeforeCancel(req *types.SfBeforeCancelRequest) (resp *types.EmptyType, err error) {
	// todo: add your logic here and delete this line

	return
}
