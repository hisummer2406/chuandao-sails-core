package platform

type PlatformManager struct {
	adapters map[string]PlatformAdapter
}

// PlatformAdapter 配送编排中心 平台调用、策略执行、流程编排
type PlatformAdapter interface {
	Inquiry()     //询价
	CreateOrder() //发单
	CancelOrder() //取消
	AddTip()      //加小费
	QueryOrder()  //查询订单
}
