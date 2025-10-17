package constants

// EventType 常量
const (
	EventOrderPush      = "ORDER_PUSH"       // 订单推送
	EventUpstreamCancel = "UPSTREAM_CANCEL"  // 上游取消
	EventUpstreamAddTip = "UPSTREAM_ADD_TIP" // 上游加小费
	EventUpstreamRetry  = "UPSTREAM_RETRY"   // 订单重推
)

// Topic 常量
const (
	TopicUpStreamPush = "upstream-push-topic"
)

// Tag 常量
const (
	TagOrderPush = "ORDER_PUSH"
)
