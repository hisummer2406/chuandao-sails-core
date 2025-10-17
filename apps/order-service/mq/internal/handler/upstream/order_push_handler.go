package upstream

import (
	"chuandao-sails-core/apps/order-service/mq/internal/logic/upstream"
	"chuandao-sails-core/apps/order-service/mq/internal/svc"
	platformEvents "chuandao-sails-core/apps/platform-gateway/pkg/platform/events"
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPushHandler struct {
	logx.Logger
	svcCtx      *svc.ServiceContext
	createLogic *upstream.OrderCreateLogic
}

func NewOrderPushHandler(svcCtx *svc.ServiceContext) *OrderPushHandler {
	return &OrderPushHandler{
		Logger:      logx.WithContext(context.Background()),
		svcCtx:      svcCtx,
		createLogic: upstream.NewOrderCreateLogic(svcCtx),
	}
}

func (h *OrderPushHandler) Handle(ctx context.Context, msg *rmq_client.MessageView) error {
	//队列中的基本信息
	messageID := msg.GetMessageId()
	topic := msg.GetTopic()
	tag := msg.GetTag()

	h.Infof("order-mp handle Received order push message, messageID: %s, topic: %s, tag: %s", messageID, topic, tag)

	//1. 解析消息体
	var orderPushEvent platformEvents.OrderPushEvent
	if err := jsonx.Unmarshal(msg.GetBody(), &orderPushEvent); err != nil {
		h.Errorf("order-mp handle push order event json unmarshal err:%s", err.Error())
		return nil
	}

	//提取TraceID
	traceID := orderPushEvent.TraceID
	if traceID == "" {
		traceID = messageID
	}
	ctx = context.WithValue(ctx, "traceID", traceID)

	//2. 调用业务逻辑创建订单
	if err := h.createLogic.OrderCreate(ctx, orderPushEvent.Data); err != nil {
		h.Errorf("order-mp handle order-push create order err:%s", err.Error())
		return fmt.Errorf("order-mp handle order-push create order err:%s", err.Error())
	}

	return nil

}
