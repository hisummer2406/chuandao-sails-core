package event

import (
	"chuandao-sails-core/apps/order-service/pkg/constants"
	"chuandao-sails-core/common/rocketmq"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type Publisher struct {
	client *rocketmq.Client
	logger logx.Logger
}

func NewPublisher(client *rocketmq.Client) *Publisher {
	return &Publisher{
		client: client,
		logger: logx.WithContext(context.Background()),
	}
}

// PublishOrderEvent 发布订单事件
func (p *Publisher) PublishOrderEvent(ctx context.Context, eventType, orderID, orderNo string, data interface{}) error {
	msg := NewOrderEventMessage(eventType, orderID, orderNo, data)

	if err := p.client.Send(ctx, msg); err != nil {
		p.logger.Errorf("Failed to publish order events: %s for order %s: %v", eventType, orderID, err)
		return fmt.Errorf("publish order events: %w", err)
	}

	p.logger.Infof("Successfully published: %s eventID: %s for orderID %s", eventType, msg.EventID, orderID)

	return nil
}

// PublishStatusChanged 发布状态变更事件
func (p *Publisher) PublishStatusChanged(ctx context.Context, orderID, orderNo, oldStatus, newStatus, reason string) error {
	data := &StatusChangedEventData{
		OrderID:    orderID,
		OldStatus:  oldStatus,
		NewStatus:  newStatus,
		ChangeTime: time.Time{},
		Reason:     reason,
	}

	return p.PublishOrderEvent(ctx, constants.EventStatusChanged, orderID, orderNo, data)
}

// PublishPenaltyCalculated 发布违约金事件
func (p *Publisher) PublishPenaltyCalculated(ctx context.Context, orderID, orderNo string, penaltyData *PenaltyEventData) error {
	return p.PublishOrderEvent(ctx, constants.EventPenaltyCalculated, orderID, orderNo, penaltyData)
}

// PublishLifecycleEvent 发布生命周期事件
func (p *Publisher) PublishLifecycleEvent(ctx context.Context, orderID, orderNo, status string, data interface{}) error {
	eventType := p.getLifecycleEventType(status)
	if eventType == "" {
		return nil
	}

	return p.PublishOrderEvent(ctx, eventType, orderID, orderNo, data)
}

// getLifecycleEventType 生命周期事件
func (p *Publisher) getLifecycleEventType(status string) string {
	switch status {
	case constants.OrderStatusCreated:
		return constants.EventOrderCreated
	case constants.OrderStatusConfirmed:
		return constants.EventOrderConfirmed
	case constants.OrderStatusAssigned:
		return constants.EventOrderAssigned
	case constants.OrderStatusPicked:
		return constants.EventOrderPicked
	case constants.OrderStatusDelivered:
		return constants.EventOrderDelivered
	case constants.OrderStatusCompleted:
		return constants.EventOrderCompleted
	case constants.OrderStatusCancelled:
		return constants.EventOrderCancelled
	default:
		return ""
	}

}
