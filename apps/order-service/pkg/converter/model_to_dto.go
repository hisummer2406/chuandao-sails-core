package converter

import "chuandao-sails-core/apps/order-service/model"

// OrderDTO 订单DTO（用于API返回）
type OrderDTO struct {
	// 基本信息
	OrderId         string `json:"order_id"`
	UpstreamOrderId string `json:"upstream_order_id"`
	Platform        string `json:"platform"`
	PlatformName    string `json:"platform_name"`
	OrderSource     string `json:"order_source"`
	OrderNum        string `json:"order_num,omitempty"`

	// 状态信息
	Status               string `json:"status"`
	StatusName           string `json:"status_name"`
	DeliveryPlatform     int    `json:"delivery_platform"`
	DeliveryPlatformName string `json:"delivery_platform_name,omitempty"`
	DeliveryOrderId      string `json:"delivery_order_id,omitempty"`

	// 地址信息
	CityName    string     `json:"city_name,omitempty"`
	CountyName  string     `json:"county_name,omitempty"`
	FromAddress AddressDTO `json:"from_address"`
	ToAddress   AddressDTO `json:"to_address"`

	// 订单类型
	SendType     int    `json:"send_type"`
	SendTypeName string `json:"send_type_name"`
	DeliveryType string `json:"delivery_type,omitempty"`

	// 价格信息
	OriginalPrice  string `json:"original_price"`
	Price          string `json:"price"`
	QuotationPrice string `json:"quotation_price"`
	DeliveryFee    int    `json:"delivery_fee"`
	TipAmount      int    `json:"tip_amount"`
	CancelFee      int    `json:"cancel_fee"`

	// 时间信息
	OrderTime     string `json:"order_time,omitempty"`
	IsSubscribe   bool   `json:"is_subscribe"`
	SubscribeTime int64  `json:"subscribe_time,omitempty"`

	// 备注
	Note string `json:"note,omitempty"`

	// 禁用配送方
	DisableDeliveryList []int `json:"disable_delivery_list,omitempty"`

	// 系统字段
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AddressDTO 地址DTO
type AddressDTO struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"` // 已脱敏
	Address   string `json:"address"`
	Detail    string `json:"detail,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

// ModelToOrderDTO Model转DTO（用于API返回）
func ModelToOrderDTO(order *model.PlatformPushOrderMaster) *OrderDTO {
	if order == nil {
		return nil
	}

	dto := &OrderDTO{
		// 基本信息
		OrderId:         order.OrderNo,
		UpstreamOrderId: order.UpstreamOrderId,
		Platform:        order.PlatformCode,
		PlatformName:    getPlatformName(order.PlatformCode),
		OrderSource:     order.OrderSource,
		OrderNum:        order.OrderNum,

		// 状态信息
		Status:               order.Status,
		StatusName:           getStatusName(order.Status),
		DeliveryPlatform:     int(order.DeliveryPlatform),
		DeliveryPlatformName: order.DeliveryPlatformName,
		DeliveryOrderId:      order.DeliveryOrderId,

		// 地址信息
		CityName:   order.CityName,
		CountyName: order.CountyName,

		// 发货地址
		FromAddress: AddressDTO{
			Name:      order.FromName,
			Phone:     maskPhone(order.FromPhone), // 脱敏
			Address:   order.FromAddress,
			Detail:    order.FromDetail,
			Latitude:  order.FromLatitude,
			Longitude: order.FromLongitude,
		},

		// 收货地址
		ToAddress: AddressDTO{
			Name:      order.ToName,
			Phone:     maskPhone(order.ToPhone), // 脱敏
			Address:   order.ToAddress,
			Detail:    order.ToDetail,
			Latitude:  order.ToLatitude,
			Longitude: order.ToLongitude,
		},

		// 订单类型
		SendType:     int(order.SendType),
		SendTypeName: getSendTypeName(int(order.SendType)),
		DeliveryType: order.DeliveryType,

		// 价格信息
		OriginalPrice:  order.OriginalPrice,
		Price:          order.Price,
		QuotationPrice: order.QuotationPrice,
		DeliveryFee:    int(order.DeliveryFee),
		TipAmount:      int(order.TipAmount),
		CancelFee:      int(order.CancelFee),

		// 时间信息
		OrderTime:     order.OrderTime,
		IsSubscribe:   order.IsSubscribe == 1,
		SubscribeTime: order.SubscribeTime,

		// 备注
		Note: order.Note,

		// 系统字段
		CreatedAt: order.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: order.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	// 解析禁用配送方列表
	if order.DisableDeliveryList != "" {
		var disableList []int
		if err := json.Unmarshal([]byte(order.DisableDeliveryList), &disableList); err == nil {
			dto.DisableDeliveryList = disableList
		}
	}

	return dto
}

// getPlatformName 获取平台名称
func getPlatformName(platform string) string {
	names := map[string]string{
		"PUU":  "UU跑腿",
		"PSF":  "顺丰同城",
		"PCHD": "船到",
	}
	if name, ok := names[platform]; ok {
		return name
	}
	return platform
}

// getStatusName 获取状态名称
func getStatusName(status string) string {
	names := map[string]string{
		"CREATED":   "已创建",
		"CONFIRMED": "已确认",
		"ASSIGNED":  "已分配",
		"PICKED":    "已取件",
		"DELIVERED": "已送达",
		"COMPLETED": "已完成",
		"CANCELLED": "已取消",
		"FAILED":    "失败",
	}
	if name, ok := names[status]; ok {
		return name
	}
	return status
}

// getSendTypeName 获取订单类型名称
func getSendTypeName(sendType int) string {
	names := map[int]string{
		0: "帮我送",
		1: "帮我买",
		2: "帮我取",
	}
	if name, ok := names[sendType]; ok {
		return name
	}
	return "未知"
}
