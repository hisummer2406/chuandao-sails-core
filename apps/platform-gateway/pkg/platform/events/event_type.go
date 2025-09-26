package events

import "time"

// StandardOrderCreateEvent 标准订单创建事件
type StandardOrderCreateEvent struct {
	// === 基本信息 ===
	OrderId         string `json:"order_id"`          // 系统内部订单号
	UpstreamOrderId string `json:"upstream_order_id"` // 上游平台订单号
	Platform        string `json:"platform"`          // 平台标识：UU/SF/CHD
	OrderSource     string `json:"order_source"`      // 订单来源
	OrderNum        string `json:"order_num"`         // 订单流水号
	CityName        string `json:"city_name"`         // 城市名称
	CountyName      string `json:"county_name"`       // 县级地名称
	AdCode          string `json:"ad_code"`           // 城市编码（高德规范）

	// === 订单类型 ===
	SendType       int    `json:"send_type"`        // 订单小类：0帮我送 1帮我买 2帮我取 UU
	IsReverseOrder bool   `json:"is_reverse_order"` // 是否帮我取订单 SF
	DeliveryType   string `json:"delivery_type"`    // 配送类型：1团送 2专送
	PushType       int    `json:"push_type"`        // 推送类型：0正常 1测试 UU

	// === 地址信息 ===
	FromAddress *AddressInfo `json:"from_address"` // 发货地址
	ToAddress   *AddressInfo `json:"to_address"`   // 收货地址

	// === 时间信息 ===
	OrderTime     string `json:"order_time"`     // 接收到订单的时间
	IsSubscribe   bool   `json:"is_subscribe"`   // 是否预约单
	SubscribeType int    `json:"subscribe_type"` // 预约类型：0实时 1预约取件 2预约送达
	SubscribeTime int64  `json:"subscribe_time"` // 预约时间戳

	// === 商品信息 ===
	GoodsInfo *GoodsInfo `json:"goods_info"` // 商品详情

	// === 价格信息 ===
	PriceInfo *PriceInfo `json:"price_info"` // 价格信息

	// === 配送选项 ===
	DeliveryOptions *DeliveryOptions `json:"delivery_options"` // 配送选项

	// === 附加信息 ===
	Note                string `json:"note"`                  // 备注
	DisableDelivery     string `json:"disable_delivery"`      // 禁用配送方
	DisableDeliveryList []int  `json:"disable_delivery_list"` // 禁用配送方ID列表

	// === 系统字段 ===
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

// AddressInfo 地址信息
type AddressInfo struct {
	Name      string `json:"name"`      // 联系人姓名
	Phone     string `json:"phone"`     // 联系电话
	Address   string `json:"address"`   // 地址
	Detail    string `json:"detail"`    // 详细地址/门牌号
	Latitude  string `json:"latitude"`  // 纬度（高德坐标）
	Longitude string `json:"longitude"` // 经度（高德坐标）
	UniqueId  string `json:"unique_id"` // 唯一ID（店铺ID/用户ID，仅发货地址）
}

// GoodsInfo 商品信息
type GoodsInfo struct {
	GoodsType   string              `json:"goods_type"`   // 商品类型
	GoodsClass  int                 `json:"goods_class"`  // 商品类型枚举
	GoodsCount  string              `json:"goods_count"`  // 商品数量
	Weight      string              `json:"weight"`       // 总重量（单位：g）
	GoodsWeight int64               `json:"goods_weight"` // 物品重量（单位：KG）
	TotalAmount string              `json:"total_amount"` // 货品总值
	ProductList []ProductDetailItem `json:"product_list"` // 商品明细列表
	ProductJson string              `json:"product_json"` // 商品列表JSON（原始）
}

// ProductDetailItem 商品明细项
type ProductDetailItem struct {
	GoodsName string `json:"goods_name"` // 商品名称
	GoodsNum  string `json:"goods_num"`  // 商品数量
	Price     string `json:"price"`      // 商品价格（元）
	Unit      string `json:"unit"`       // 商品规格
	Weight    string `json:"weight"`     // 商品重量（克）
	GoodsImg  string `json:"goods_img"`  // 商品图片
	Amount    string `json:"amount"`     // 金额（分）
}

// PriceInfo 价格信息
type PriceInfo struct {
	OriginalPrice  string `json:"original_price"`  // 商品原价（元）
	Price          string `json:"price"`           // 实际支付金额（元）
	QuotationPrice string `json:"quotation_price"` // 报价（元）
	PostPrice      string `json:"post_price"`      // 发单价格（分，包含小费）
	OrderTip       string `json:"order_tip"`       // 订单小费（分）
	GoodsPrice     int    `json:"goods_price"`     // 商品保价价格（元）
	PriceDistance  string `json:"price_distance"`  // 计价距离（米）
}

// DeliveryOptions 配送选项
type DeliveryOptions struct {
	IsInsured      bool   `json:"is_insured"`       // 是否保价
	IsInsurance    bool   `json:"is_insurance"`     // 是否保价（兼容字段）CHD
	IsPersonDirect bool   `json:"is_person_direct"` // 是否专人直送 SF
	IsDeliveryImg  bool   `json:"is_delivery_img"`  // 是否需要妥投照片 UU
	IsCityShop     bool   `json:"is_city_shop"`     // 是否城市型店铺 SF
	WeatherLevel   string `json:"weather_level"`    // 天气等级 SF
	SfStoreId      string `json:"sf_store_id"`      // 顺丰门店ID 小镇外卖兼容字段
}
