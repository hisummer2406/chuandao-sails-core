package constants

const (
	//订单生命周期事件
	EventOrderCreated   = "order.created"   //订单已创建
	EventOrderConfirmed = "order.confirmed" //订单已确认
	EventOrderAssigned  = "order.assigned"  //订单已分配配送员
	EventOrderPicked    = "order.picked"    //订单已取件
	EventOrderDelivered = "order.delivered" //订单已送单
	EventOrderCompleted = "order.completed" //订单已完成
	EventOrderCancelled = "order.cancelled" //订单已取消

	//状态变更时间
	EventStatusChanged = "status.changed" //状态变更

	//违约事件
	EventPenaltyCalculated = "penalty.calculated" //违约金已计算
	EventPenaltySettled    = "penalty.settled"    //违约金已结算

	//异常事件
	EventOrderTimeout = "order.timeout" //订单超时
	EventOrderFailed  = "order.failed"  //订单失败
)
