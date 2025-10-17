package events

// BaseEvent 统一事件基础结构
type BaseEvent struct {
	EventID   string `json:"event_id"`   // 事件唯一ID
	EventType string `json:"event_type"` // 事件类型
	Source    string `json:"source"`     // 事件源服务
	Timestamp int64  `json:"timestamp"`  // 事件时间戳
	TraceID   string `json:"trace_id"`   // 链路追踪ID
	Version   string `json:"version"`    // 事件版本
}
