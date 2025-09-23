package client

import "context"

type DeliveryClient interface {
	StartDispatch(ctx context.Context, req *DeliveryDispatchRequest) error
}

type DeliveryDispatchRequest struct {
	OrderID string `json:"order_id"`
}
