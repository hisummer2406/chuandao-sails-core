package model

import (
	"chuandao-sails-core/demo/restful/user/internal/svc"
	"context"
	"encoding/json"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

type Consumer struct {
	svcCtx *svc.ServiceContext
}

func NewConsumer(svcCtx *svc.ServiceContext) *Consumer {
	return &Consumer{
		svcCtx: svcCtx,
	}
}

// Start 启动消费者
func (c *Consumer) Start() error {
	//订阅订单主题
	if err := c.svcCtx.RocketMQ.Subscribe("order_topic", c.HandlerOrderMessage); err != nil {
		return err
	}

	//订阅支付主题
	if err := c.svcCtx.RocketMQ.Subscribe("payment_topic", c.HandlerPaymentMessage); err != nil {
		return err
	}

	return c.svcCtx.RocketMQ.StartConsumer()
}

// HandlerOrderMessage 处理订单消息
func (c *Consumer) HandlerOrderMessage(ctx context.Context, msg *rmq_client.MessageView) error {
	logx.Infof("rocketmq receive order message: id=%s, tag=%s", msg.GetMessageId(), *msg.GetTag())

	var order OrderMessage
	if err := json.Unmarshal(msg.GetBody(), &order); err != nil {
		logx.Errorf("rocketmq unmarshal order message failed, err=%v", err)
		return err
	}

	//根据标签处理不同事件
	switch *msg.GetTag() {
	case "created":
		//处理订单创建事件
	case "paid":
		//处理订单支付事件
	case "canceled":
	//处理订单取消事件
	default:
		logx.Infof("rocketmq receive order message: tag=%s not support", *msg.GetTag())
		return nil
	}

	return nil
}

// HandlerPaymentMessage 处理支付消息
func (c *Consumer) HandlerPaymentMessage(ctx context.Context, msg *rmq_client.MessageView) error {
	logx.Infof("rocketmq receive payment message: id=%s", msg.GetMessageId())

	var payment PaymentMessage
	if err := json.Unmarshal(msg.GetBody(), &payment); err != nil {
		logx.Errorf("rocketmq unmarshal payment message failed, err=%v", err)
		return err
	}

	//根据标签处理不同事件
	switch *msg.GetTag() {
	case "success":
		//处理支付成功事件
	case "failed":
		//处理支付失败事件
	default:
		logx.Infof("rocketmq receive payment message: tag=%s not support", *msg.GetTag())
		return nil
	}

	return c.handlePaymentSuccess(ctx, &payment)
}
