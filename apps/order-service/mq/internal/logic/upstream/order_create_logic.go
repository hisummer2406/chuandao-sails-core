package upstream

import (
	"chuandao-sails-core/apps/order-service/mq/internal/svc"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/events"
	"chuandao-sails-core/common/snowflake"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		Logger: logx.WithContext(context.Background()),
		svcCtx: svcCtx,
	}
}

func (l *OrderCreateLogic) OrderCreate(ctx context.Context, event *events.StandardOrderCreateEvent) error {
	l.ctx = ctx
	traceID := ctx.Value("traceID").(string)
	l.Infof("[TraceID: %s] order-mp Start creating order, upstreamOrderId: %s, platform: %s", traceID, event.UpstreamOrderId, event.PlatformCode)

	//1. 幂等性检查
	existingOrder, err := l.svcCtx.OrderModel.FindOneByUpstreamOrderId(ctx, event.UpstreamOrderId)
	if err != nil {
		l.Errorf("[TraceID: %s] order-mp OrderCreate existingOrder failed: %v", traceID, err)
		return fmt.Errorf("%s OrderCreate existingOrder failed: %w", event.UpstreamOrderId, err)
	}
	// 订单已经创建
	if existingOrder != nil {
		l.Infof("[TraceID: %s] order-mp Order already exists , orderId: %s skip creating", traceID, event.UpstreamOrderId)
		return nil
	}

	//2. 生成单号
	orderNo, err := snowflake.GenerateOrderNoWithPrefix("ORD")
	if err != nil {
		l.Errorf("[TraceID: %s] order-mp Order GenerateOrderNoWithPrefix failed: %v", traceID, err)
		return fmt.Errorf("%s OrderCreate GenerateOrderNoWithPrefix failed: %w", event.UpstreamOrderId, err)
	}
	event.OrderNo = orderNo

	//3.数据转换

	return nil
}
