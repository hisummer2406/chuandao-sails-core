package event

import (
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/jsonx"
	"time"
)

// 订单创建事件
var topic string = "order-create-topic"

type OrderEventMessage struct {
	EventID   string                 `json:"event_id"`
	EventType string                 `json:"event_type"`
	OrderID   string                 `json:"order_id"`
	OrderNo   string                 `json:"order_no"`
	Source    string                 `json:"source"`
	Timestamp time.Time              `json:"timestamp"`
	Data      interface{}            `json:"data"`
	Metadata  map[string]interface{} `json:"metadata"`
}

func (o *OrderEventMessage) GetTopic() string {
	return topic
}

func (o *OrderEventMessage) GetTag() string {
	return o.EventType
}

func (o *OrderEventMessage) GetKeys() []string {
	//消息索引是云消息队列 RocketMQ 版提供的面向消息的索引属性。通过设置的消息索引可以快速查找到对应的消息内容。
	return []string{o.OrderID, o.EventID}
}

func (o *OrderEventMessage) ToBytes() ([]byte, error) {
	return jsonx.Marshal(o)
}

// NewOrderEventMessage 创建订单事件消息
func NewOrderEventMessage(eventType, orderID, orderNo string, data interface{}) *OrderEventMessage {
	return &OrderEventMessage{
		EventID:   uuid.New().String(),
		EventType: eventType,
		OrderID:   orderID,
		OrderNo:   orderNo,
		Source:    topic,
		Data:      data,
		Metadata: map[string]interface{}{
			"service": topic,
			"version": "1.0",
		},
	}
}
