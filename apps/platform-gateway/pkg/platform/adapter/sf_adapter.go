package adapter

import (
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"chuandao-sails-core/apps/platform-gateway/pkg/constants"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/events"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/utils"
	"chuandao-sails-core/common/snowflake"
	"github.com/zeromicro/go-zero/core/jsonx"
	"time"
)

// SFAdapter 顺丰平台适配器
type SFAdapter struct{}

func (a *SFAdapter) TransformToStandardOrder(req *types.SfCreateRequest) (*events.StandardOrderCreateEvent, error) {
	// 解析商品信息
	var productList []events.ProductDetailItem
	if req.Goods != "" {
		jsonx.Unmarshal([]byte(req.Goods), &productList)
	}

	orderId, err := snowflake.GenerateOrderId()
	if err != nil {
		return nil, err
	}

	// 字符串转布尔
	isSubscribe := req.IsSubscribe == "1"
	isInsured := req.IsInsured == "1"
	isPersonDirect := req.IsPersonDirect == "1"
	isCityShop := req.IsCityShop == "1"
	isReverseOrder := req.IsReverseOrder == "1"

	return &events.StandardOrderCreateEvent{
		// 基本信息
		OrderId:         orderId,
		UpstreamOrderId: req.OrderNo,
		Platform:        constants.PLATFORM_SF,
		OrderSource:     req.Source,
		OrderNum:        req.OrderSn,
		CityName:        "", // 通过AdCode获取
		AdCode:          req.AdCode,

		// 订单类型
		SendType:       utils.GetSendTypeFromSF(isReverseOrder),
		IsReverseOrder: isReverseOrder,

		// 地址信息
		FromAddress: &events.AddressInfo{
			Name:      req.FromName,
			Phone:     req.FromPhone,
			Address:   req.FromAddress,
			Detail:    req.FromAddressDetail,
			Latitude:  req.FromLatitude,
			Longitude: req.FromLongitude,
			UniqueId:  req.FromId,
		},
		ToAddress: &events.AddressInfo{
			Name:      req.ToName,
			Phone:     req.ToPhone,
			Address:   req.ToAddress,
			Detail:    req.ToAddressDetail,
			Latitude:  req.ToLatitude,
			Longitude: req.ToLongitude,
		},

		// 时间信息
		OrderTime:     req.OrderTime,
		IsSubscribe:   isSubscribe,
		SubscribeTime: utils.ParseMilliseconds(req.SubscribeTime),

		// 商品信息
		GoodsInfo: &events.GoodsInfo{
			GoodsType:   req.GoodType,
			GoodsCount:  req.GoodsCount,
			Weight:      req.Weight,
			TotalAmount: req.TotalAmount,
			ProductList: productList,
			ProductJson: req.Goods,
		},

		// 价格信息
		PriceInfo: &events.PriceInfo{
			PostPrice:     req.PostPrice,
			OrderTip:      req.OrderTip,
			PriceDistance: req.PriceDistance,
		},

		// 配送选项
		DeliveryOptions: &events.DeliveryOptions{
			IsInsured:      isInsured,
			IsPersonDirect: isPersonDirect,
			IsCityShop:     isCityShop,
			WeatherLevel:   req.WeatherLevel,
		},

		// 附加信息
		Note:      req.OrderRemark,
		SfStoreId: "", // 在请求中未直接体现

		// 系统字段
		CreatedAt: time.Now(),
	}, nil
}
