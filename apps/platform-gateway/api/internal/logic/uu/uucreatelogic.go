package uu

import (
	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"chuandao-sails-core/apps/platform-gateway/pkg/constants"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/adapter"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/events"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/utils"
	"chuandao-sails-core/common/response"
	"chuandao-sails-core/common/tools"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"regexp"
	"strings"
)

type UuCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUuCreateLogic 接收订单
func NewUuCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UuCreateLogic {
	return &UuCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UuCreateLogic) UuCreate(req *types.UUCreateOrderRequest) (resp *types.EmptyType, err error) {

	//数据格式标准化
	factory := &adapter.AdapterFactory{}
	uuAdapter := factory.GetAdapter(constants.PLATFORM_UU)
	//转换为标准订单事件
	standardOrder, err := uuAdapter.TransformToStandardOrder(req)
	if err != nil {
		return nil, response.NewBusinessErrorWithCtx(l.ctx, http.StatusBadRequest, fmt.Sprintf("UuCreate transfer error : %s", err.Error()))
	}

	//数据校验和优化
	l.validateRequest(standardOrder, req)

	// TODO: 发送订单事件到消息队列
	// 这里应该发送到 order-create-topic，由订单服务接收处理
	// err = l.svcCtx.OrderProducer.SendOrderCreateEvent(standardOrder)
	// if err != nil {
	//     l.Logger.Errorf("发送订单创建事件失败: %v", err)
	//     return nil, response.NewInternalServerErrorWithCtx(l.ctx, "订单创建失败")
	// }

	return &types.EmptyType{}, nil
}

// validateRequest 数据校验和优化
func (l *UuCreateLogic) validateRequest(standardOrder *events.StandardOrderCreateEvent, req *types.UUCreateOrderRequest) error {
	//解析禁用配送方
	standardOrder.DisableDeliveryList = utils.ParseDisableDelivery(req.DisableDelivery)
	//订单来源
	standardOrder.OrderSource = utils.ValidateOrderSource(req.OrderSource)
	//商品类型
	standardOrder.GoodsInfo.GoodsClass = utils.ValidateGoodsClass(req.GoodsClass)
	//手机号验证
	l.processReceiverPhone(standardOrder)
	//取餐号
	standardOrder.OrderNum = tools.ProcessPickNo(standardOrder.OrderNum)
	//TODO 默认价格、校验类目 都在MQ接受中处理
	return nil
}

// processReceiverPhone 隐私号处理
func (l *UuCreateLogic) processReceiverPhone(standardOrder *events.StandardOrderCreateEvent) {
	phone := standardOrder.ToAddress.Phone

	// 检查是否包含中文或英文字符
	hasChinese := regexp.MustCompile(`[\x{4e00}-\x{9fa5}]`).MatchString(phone)
	hasEnglish := regexp.MustCompile(`[a-zA-Z]`).MatchString(phone)

	if hasChinese || hasEnglish {
		// 尝试提取手机号格式: 11位数字_3-4位分机号
		re := regexp.MustCompile(`(\d{11}[_ ,.-]\d{3,4})`)
		matches := re.FindStringSubmatch(phone)
		if len(matches) > 1 {
			phone = matches[1]
		} else {
			// 截取前50个字符
			runes := []rune(phone)
			if len(runes) > 50 {
				phone = string(runes[:50])
			}
		}
	}

	// 处理隐私号: 手机号,分机号 或 手机号_分机号
	phone = strings.ReplaceAll(phone, "_", ",")
	parts := strings.Split(phone, ",")

	// 提取手机号后4位
	var suffix string
	if len(parts) > 0 && len(parts[0]) >= 4 {
		suffix = parts[0][len(parts[0])-4:]
	}

	// 如果有分机号，使用分机号
	if len(parts) > 1 {
		suffix = parts[1]
	}

	// 验证后缀长度不超过4位
	if len(suffix) > 4 {
		suffix = "0"
	}
	if suffix == "" {
		suffix = "0"
	}

	// 更新地址信息
	standardOrder.ToAddress.Phone = phone

}
