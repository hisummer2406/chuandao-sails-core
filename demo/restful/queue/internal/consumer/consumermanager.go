package consumer

import (
	"chuandao-sails-core/demo/restful/queue/internal/logic"
	"chuandao-sails-core/demo/restful/queue/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ConsumerManager struct {
	svcCtx    *svc.ServiceContext
	processor *logic.MessageProcessor
}

func NewConsumerManager(svcCtx *svc.ServiceContext) *ConsumerManager {
	return &ConsumerManager{
		svcCtx:    svcCtx,
		processor: logic.NewMessageProcessor(svcCtx),
	}
}

// Start 启动所有消费者
func (cm *ConsumerManager) Start() error {
	logx.Info("Starting message consumers...")

	// 订阅订单主题
	if err := cm.svcCtx.RocketMQ.Subscribe("order_topic", cm.processor.ProcessOrderMessage); err != nil {
		return err
	}
	logx.Info("Subscribed to order_topic")

	// 订阅支付主题
	if err := cm.svcCtx.RocketMQ.Subscribe("payment_topic", cm.processor.ProcessPaymentMessage); err != nil {
		return err
	}
	logx.Info("Subscribed to payment_topic")

	// 启动消费者
	if err := cm.svcCtx.RocketMQ.StartConsumer(); err != nil {
		return err
	}

	logx.Info("All message consumers started successfully")
	return nil
}
