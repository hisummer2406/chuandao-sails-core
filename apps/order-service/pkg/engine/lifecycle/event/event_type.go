package event

import "time"

// 基础事件结构
type BaseEvent struct {
	EventID   string                 `json:"event_id"`   //事件唯一ID
	EventType string                 `json:"event_type"` //事件类型
	Source    string                 `json:"source"`     //事件源
	Timestamp int64                  `json:"timestamp"`  //事件时间
	Version   string                 `json:"version"`    //时间版本
	Metadata  map[string]interface{} `json:"metadata"`   //元数据
}

// 订单事件
type OrderEvent struct {
	BaseEvent
	OrderID         string      `json:"order_id"`          //订单ID
	OrderNo         string      `json:"order_no"`          //订单号
	PlatformOrderNo string      `json:"platform_order_no"` //平台订单号
	CustomerID      string      `json:"customer_id"`       //客户ID
	Data            interface{} `json:"data"`              //事件数据
}

// 状态变更事件数据
type StatusChangedEventData struct {
	OrderID    string                 `json:"order_id"`
	OldStatus  string                 `json:"old_status"`  //原状态
	NewStatus  string                 `json:"new_status"`  //新状态
	ChangeTime time.Time              `json:"change_time"` //变更时间
	Reason     string                 `json:"reason"`      //变更原因
	OperatorID string                 `json:"operator_id"` //操作员ID
	Extra      map[string]interface{} `json:"extra"`       //额外信息
}

// 违约金事件数据
type PenaltyEventData struct {
	OrderID       string    `json:"order_id"`
	PenaltyType   string    `json:"penalty_type"`   //违约类型
	PenaltyAmount float64   `json:"penalty_amount"` //违约金额
	PenaltyReason string    `json:"penalty_reason"` //违约原因
	CalculateTime time.Time `json:"calculate_time"` //计算时间
	Rules         []string  `json:"rules"`          //应用的规则
}
