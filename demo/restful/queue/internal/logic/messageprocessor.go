package logic

import (
	"chuandao-sails-core/demo/restful/queue/internal/model"
	"chuandao-sails-core/demo/restful/queue/internal/svc"
	"context"
	"encoding/json"

	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

type MessageProcessor struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageProcessor(svcCtx *svc.ServiceContext) *MessageProcessor {
	return &MessageProcessor{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

// ProcessOrderMessage 处理订单消息
func (p *MessageProcessor) ProcessOrderMessage(ctx context.Context, msg *golang.MessageView) error {
	p.Infof("Processing order message: id=%s", msg.GetMessageId())

	var order model.OrderMessage
	if err := json.Unmarshal(msg.GetBody(), &order); err != nil {
		p.Errorf("Failed to parse order message: %v", err)
		return err
	}

	// 根据订单状态处理
	switch order.Status {
	case "CREATED":
		return p.handleOrderCreated(ctx, &order)
	case "PAID":
		return p.handleOrderPaid(ctx, &order)
	case "CANCELLED":
		return p.handleOrderCancelled(ctx, &order)
	default:
		p.Infof("Unknown order status: %s", order.Status)
		return nil
	}
}

// ProcessPaymentMessage 处理支付消息
func (p *MessageProcessor) ProcessPaymentMessage(ctx context.Context, msg *golang.MessageView) error {
	p.Infof("Processing payment message: id=%s", msg.GetMessageId())

	var payment model.PaymentMessage
	if err := json.Unmarshal(msg.GetBody(), &payment); err != nil {
		p.Errorf("Failed to parse payment message: %v", err)
		return err
	}

	return p.handlePaymentSuccess(ctx, &payment)
}

func (p *MessageProcessor) handleOrderCreated(ctx context.Context, order *model.OrderMessage) error {
	p.Infof("Handling order created: %s", order.OrderID)
	// TODO: 实现订单创建处理逻辑
	return nil
}

func (p *MessageProcessor) handleOrderPaid(ctx context.Context, order *model.OrderMessage) error {
	p.Infof("Handling order paid: %s", order.OrderID)
	// TODO: 实现订单支付处理逻辑
	return nil
}

func (p *MessageProcessor) handleOrderCancelled(ctx context.Context, order *model.OrderMessage) error {
	p.Infof("Handling order cancelled: %s", order.OrderID)
	// TODO: 实现订单取消处理逻辑
	return nil
}

func (p *MessageProcessor) handlePaymentSuccess(ctx context.Context, payment *model.PaymentMessage) error {
	p.Infof("Handling payment success: %s", payment.PaymentID)
	// TODO: 实现支付成功处理逻辑
	return nil
}
