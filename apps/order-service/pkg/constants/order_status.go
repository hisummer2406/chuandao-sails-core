package constants

const (
	OrderStatusCreated   = "CREATED"   // 已创建
	OrderStatusConfirmed = "CONFIRMED" // 已确认
	OrderStatusAssigned  = "ASSIGNED"  // 已分配配送员
	OrderStatusPicked    = "PICKED"    // 已取件
	OrderStatusDelivered = "DELIVERED" // 已送达
	OrderStatusCompleted = "COMPLETED" // 已完成
	OrderStatusCancelled = "CANCELLED" // 已取消
	OrderStatusFailed    = "FAILED"    //失败
)

// StateTransitions 状态转换规则定义
var StateTransitions = map[string][]string{
	OrderStatusCreated:   {OrderStatusConfirmed, OrderStatusCancelled, OrderStatusFailed},
	OrderStatusConfirmed: {OrderStatusAssigned, OrderStatusCancelled},
	OrderStatusAssigned:  {OrderStatusPicked, OrderStatusCreated},
	OrderStatusPicked:    {OrderStatusDelivered},
	OrderStatusDelivered: {OrderStatusCompleted},
	OrderStatusCompleted: {}, //终态
	OrderStatusCancelled: {}, //终态
}
