package constants

// 订单来源枚举
const (
	ORDER_SOURCE_MT    = "MT"    // 美团
	ORDER_SOURCE_SG    = "SG"    // 闪购
	ORDER_SOURCE_EL    = "EL"    // 饿了么
	ORDER_SOURCE_ELB   = "ELB"   // 饿了么零售
	ORDER_SOURCE_DY    = "DY"    // 抖音
	ORDER_SOURCE_DYD   = "DYD"   // 抖音小时达
	ORDER_SOURCE_JDDJ  = "JDDJ"  // 京东到家
	ORDER_SOURCE_MINI  = "MINI"  // 小程序
	ORDER_SOURCE_XZWM  = "XZWM"  // 小镇外卖
	ORDER_SOURCE_OTHER = "OTHER" // 其他
)

// 订单来源名称映射
var OrderSourceNames = map[string]string{
	ORDER_SOURCE_MT:    "美团",
	ORDER_SOURCE_SG:    "闪购",
	ORDER_SOURCE_EL:    "饿了么",
	ORDER_SOURCE_ELB:   "饿了么零售",
	ORDER_SOURCE_DY:    "抖音",
	ORDER_SOURCE_DYD:   "抖音小时达",
	ORDER_SOURCE_JDDJ:  "京东到家",
	ORDER_SOURCE_MINI:  "小程序",
	ORDER_SOURCE_XZWM:  "小镇外卖",
	ORDER_SOURCE_OTHER: "其他",
}
