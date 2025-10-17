package converter

import (
	"chuandao-sails-core/apps/order-service/model"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/events"
	"github.com/zeromicro/go-zero/core/jsonx"
)

func EventToOrderModel(event *events.StandardOrderCreateEvent) *model.PlatformPushOrderMaster {
	orderModel := &model.PlatformPushOrderMaster{
		//订单基本信息
		OrderNo:         event.OrderNo,
		UpstreamOrderId: event.UpstreamOrderId,
		PlatformCode:    event.PlatformCode,
		OrderSource:     event.OrderSource,
		OrderNum:        event.OrderNum,
		CityName:        event.CityName,
		CountyName:      event.CountyName,
		AdCode:          event.AdCode, // TODO 未处理

		//订单类型
		SendType:       event.SendType,
		IsReverseOrder: boolToInt64(event.IsReverseOrder),
		DeliveryType:   event.DeliveryType,
		PushType:       event.PushType,

		//地址
		FromAddress: marshalToJSON(event.FromAddress),
		ToAddress:   marshalToJSON(event.ToAddress),

		//时间
		OrderTime:     event.OrderTime,
		IsSubscribe:   boolToInt64(event.IsSubscribe),
		SubscribeType: event.SubscribeType,
		SubscribeTime: event.SubscribeTime,

		//商品
		GoodsInfo: marshalToJSON(event.GoodsInfo),
		//价格
		PriceInfo: marshalToJSON(event.PriceInfo),
		//配送选项
		DeliveryOptions: marshalToJSON(event.DeliveryOptions),
		//附加
		Note:            event.Note,
		DisableDelivery: event.DisableDelivery,
	}

	return orderModel
}

// toMarshal 转json
func marshalToJSON(v interface{}) string {
	if v == nil {
		return "{}"
	}
	data, err := jsonx.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(data)
}

// boolToInt64 bool 转 int
func boolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
