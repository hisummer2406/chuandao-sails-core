package constants

// 数据类型常量
const (
	PLATFORM_UU  = "PUU"  // UU格式
	PLATFORM_SF  = "PSF"  // 顺丰格式
	PLATFORM_CHD = "PCHD" // 船到格式
)

// 订单类型常量
const (
	SEND_TYPE_HELP_SEND = 0 // 帮我送
	SEND_TYPE_HELP_BUY  = 1 // 帮我买
	SEND_TYPE_HELP_GET  = 2 // 帮我取
)

// 订单预约类型常量
const (
	SUBSCRIBE_TYPE_REALTIME     = 0 // 实时单
	SUBSCRIBE_TYPE_PICKUP_TIME  = 1 // 预约取件时间
	SUBSCRIBE_TYPE_DELIVER_TIME = 2 // 预约送达时间
)

// 配送类型常量
const (
	DELIVERY_TYPE_GROUP     = 1 // 团送
	DELIVERY_TYPE_EXCLUSIVE = 2 // 专送
)
