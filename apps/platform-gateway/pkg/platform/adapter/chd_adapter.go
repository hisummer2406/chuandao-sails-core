package adapter

import (
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"chuandao-sails-core/apps/platform-gateway/pkg/constants"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/events"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/utils"
	"chuandao-sails-core/common/snowflake"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"strconv"
	"time"
)

// CHDAdapter 船到平台适配器
type CHDAdapter struct{}

func (a *CHDAdapter) TransformToStandardOrder(data interface{}) (*events.StandardOrderCreateEvent, error) {
	req, ok := data.(*types.ChdCreateRequest)
	if !ok {
		return nil, fmt.Errorf("invalid request type for CHDAdapter In TransformToStandardOrder")
	}
	// 解析商品明细
	var productList []events.ProductDetailItem
	if req.ProductDetail != "" {
		jsonx.Unmarshal([]byte(req.ProductDetail), &productList)
	}

	orderNo, err := snowflake.GenerateOrderNoWithPrefix(constants.PLATFORM_CHD)
	if err != nil {
		return nil, err
	}

	// 字符串转整数
	subscribeType, _ := strconv.Atoi(req.SubscribeType)

	return &events.StandardOrderCreateEvent{
		// 基本信息
		OrderNo:         orderNo,
		UpstreamOrderId: req.DeliveryId,
		PlatformCode:    constants.PLATFORM_CHD,
		OrderSource:     req.OrderSource,
		OrderNum:        req.OrderNum,
		CityName:        req.CityName,
		CountyName:      req.CountyName,

		// 订单类型
		SendType:     req.SendType,
		DeliveryType: req.DeliveryType,

		// 地址信息
		FromAddress: &events.AddressInfo{
			Name:      req.Sender,
			Phone:     req.SenderPhone,
			Address:   req.FromAddress,
			Detail:    req.FromUserNote,
			Latitude:  req.FromLat,
			Longitude: req.FromLng,
		},
		ToAddress: &events.AddressInfo{
			Name:      req.Receiver,
			Phone:     req.ReceiverPhone,
			Address:   req.ToAddress,
			Detail:    req.ToUserNote,
			Latitude:  req.ToLat,
			Longitude: req.ToLng,
		},

		// 时间信息
		IsSubscribe:   subscribeType > 0,
		SubscribeType: subscribeType,
		SubscribeTime: utils.ParseTimestamp(req.SubscribeTime),

		// 商品信息
		GoodsInfo: &events.GoodsInfo{
			GoodsClass:  req.GoodsClass,
			GoodsWeight: utils.ParseTimestamp(req.GoodsWeight),
			ProductList: productList,
			ProductJson: req.ProductDetail,
		},

		// 价格信息
		PriceInfo: &events.PriceInfo{
			OriginalPrice:  req.OriginalPrice,
			Price:          req.Price,
			QuotationPrice: req.QuotationPrice,
		},

		// 配送选项
		DeliveryOptions: &events.DeliveryOptions{
			IsInsurance: req.IsInsurance == 1,
			SfStoreId:   req.SfStoreId, //小镇外卖兼容字段 不独立建字段，写入拓展字段
		},

		// 附加信息
		Note:            req.Note,
		DisableDelivery: req.DisableDelivery,

		// 系统字段
		CreatedAt: time.Now(),
	}, nil
}
