package adapter

import (
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"chuandao-sails-core/apps/platform-gateway/pkg/constants"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/events"
	"chuandao-sails-core/common/snowflake"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"time"
)

type UUAdapter struct{}

func (a *UUAdapter) TransformToStandardOrder(data interface{}) (*events.StandardOrderCreateEvent, error) {
	req, ok := data.(*types.UUCreateOrderRequest)
	if !ok {
		return nil, fmt.Errorf("invalid request type for UUAdapter In TransformToStandardOrder")
	}

	// 解析商品明细
	var productList []events.ProductDetailItem
	if req.ProductDetail != "" {
		jsonx.Unmarshal([]byte(req.ProductDetail), &productList)
	}

	orderNo, err := snowflake.GenerateOrderNoWithPrefix(constants.PLATFORM_UU)
	if err != nil {
		return nil, err
	}

	return &events.StandardOrderCreateEvent{
		// 基本信息
		OrderNo:         orderNo,
		UpstreamOrderId: req.DeliveryId,
		PlatformCode:    constants.PLATFORM_UU,
		OrderSource:     req.OrderSource,
		OrderNum:        req.OrderNum,
		CityName:        req.CityName,
		CountyName:      req.CountyName,
		AdCode:          "", // UU未提供

		// 订单类型
		SendType:       req.SendType,      //订单小类 0帮我送 1帮我买 2帮我取
		DeliveryType:   req.DeliveryType,  //1 团送 2 专送
		IsReverseOrder: req.SendType == 2, // 帮我取
		PushType:       int(req.PushType), // 0 正常 1 测试

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
		OrderTime:     time.Now().Format("2006-01-02 15:04:05"),
		IsSubscribe:   req.SubscribeType > 0,
		SubscribeType: int(req.SubscribeType),
		SubscribeTime: req.SubscribeTime,

		// 商品信息
		GoodsInfo: &events.GoodsInfo{
			GoodsType:   req.GoodsType,
			GoodsClass:  req.GoodsClass,
			GoodsWeight: req.GoodsWeight,
			ProductList: productList,
			ProductJson: req.ProductDetail,
		},

		// 价格信息
		PriceInfo: &events.PriceInfo{
			OriginalPrice:  req.OriginalPrice,
			Price:          req.Price,
			QuotationPrice: req.QuotationPrice,
			GoodsPrice:     req.GoodsPrice,
		},

		// 配送选项
		DeliveryOptions: &events.DeliveryOptions{
			IsDeliveryImg: req.IsDeliveryImg == 1,
		},

		// 附加信息
		Note:            req.Note,
		DisableDelivery: req.DisableDelivery,

		// 系统字段
		CreatedAt: time.Now(),
	}, nil
}
