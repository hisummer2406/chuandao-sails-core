package events

import (
	"chuandao-sails-core/apps/platform-gateway/pkg/constants"
	"chuandao-sails-core/common/events"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/jsonx"
	"time"
)

type OrderPushEvent struct {
	events.BaseEvent
	Metadata map[string]interface{}    `json:"metadata"` // 元数据
	Data     *StandardOrderCreateEvent `json:"data"`     //业务数据
}

// NewOrderPushEvent 创建订单推送事件
func NewOrderPushEvent(orderData *StandardOrderCreateEvent, traceID string) *OrderPushEvent {
	return &OrderPushEvent{
		BaseEvent: events.BaseEvent{
			EventID:   uuid.New().String(),
			EventType: constants.EventOrderPush,
			Source:    "platform-gateway",
			Timestamp: time.Now().Unix(),
			TraceID:   traceID,
			Version:   "1.0.0",
		},
		Data: orderData,
	}
}

// ============ 实现 common/rocketmq BusinessMessage 接口 ============

// ToBytes 序列化为字节数组
func (e *OrderPushEvent) ToBytes() ([]byte, error) {
	return jsonx.Marshal(e)
}

// GetTopic 获取topic
func (e *OrderPushEvent) GetTopic() string {
	return constants.TopicUpStreamPush
}

// GetTag 获取tag
func (e *OrderPushEvent) GetTag() string {
	return constants.TagOrderPush
}

// GetKeys 获取消息Keys（用于消息索引和追踪）
func (e *OrderPushEvent) GetKeys() []string {
	if e.Data != nil {
		return []string{
			e.Data.UpstreamOrderId,
			e.Data.PlatformCode,
			e.EventID,
		}
	}
	return []string{e.EventID}
}
