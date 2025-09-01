package v1logic

import (
	"context"

	"chuandao-sails-core/apps/delivery-strategy/rpc/internal/svc"
	"chuandao-sails-core/apps/delivery-strategy/rpc/strategy"

	"github.com/zeromicro/go-zero/core/logx"
)

type CalculateQuoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCalculateQuoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalculateQuoteLogic {
	return &CalculateQuoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新城市特殊配置
func (l *CalculateQuoteLogic) CalculateQuote(in *strategy.CalculateQuoteReq) (*strategy.CalculateQuoteResp, error) {
	// todo: add your logic here and delete this line

	return &strategy.CalculateQuoteResp{}, nil
}
