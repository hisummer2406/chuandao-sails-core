package processor

import (
	"chuandao-sails-core/apps/order-service/pkg/client"
	"chuandao-sails-core/apps/order-service/pkg/constants"
	"chuandao-sails-core/apps/order-service/pkg/engine/lifecycle/event"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreatedProcessor struct {
	deliveryClient client.DeliveryClient
	logger         logx.Logger
}

func NewOrderCreatedHandler(deliveryClient client.DeliveryClient) *OrderCreatedProcessor {
	return &OrderCreatedProcessor{
		deliveryClient: deliveryClient,
		logger:         logx.WithContext(context.Background()),
	}
}

func (p *OrderCreatedProcessor) EventTypes() []string {
	return []string{constants.EventOrderCreated}
}

func (p *OrderCreatedProcessor) Handle(ctx context.Context, msg *event.OrderEventMessage) error {
	p.logger.Infof("Processing order event message: %v", msg)

	err := p.deliveryClient.StartDispatch(ctx, &client.DeliveryDispatchRequest{
		OrderID: msg.OrderID,
	})

	if err != nil {
		return fmt.Errorf("error starting delivery dispatch: %w", err)
	}

	return nil
}
