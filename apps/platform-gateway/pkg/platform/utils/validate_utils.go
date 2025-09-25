package utils

import (
	"chuandao-sails-core/apps/platform-gateway/pkg/constants"
	"strconv"
	"strings"
)

// 验证订单来源是否有效
func ValidateOrderSource(source string) string {
	if _, exists := constants.OrderSourceNames[source]; exists {
		return source
	}
	return constants.ORDER_SOURCE_OTHER // 默认其他
}

// 验证商品类型是否有效
func ValidateGoodsClass(class int) int {
	if _, exists := constants.GoodsClassNames[class]; exists {
		return class
	}
	return constants.GOODS_CLASS_OTHER // 默认其他
}

// 解析禁用配送方字符串为数组
func ParseDisableDelivery(disable string) []int {
	if disable == "" {
		return nil
	}

	var platforms []int
	parts := strings.Split(disable, ",")
	for _, part := range parts {
		if platform, err := strconv.Atoi(strings.TrimSpace(part)); err == nil {
			if _, exists := constants.DeliveryPlatformNames[platform]; exists {
				platforms = append(platforms, platform)
			}
		}
	}
	return platforms
}

// 检查是否禁用某个配送平台
func IsDeliveryDisabled(disableList []int, platformId int) bool {
	for _, id := range disableList {
		if id == platformId {
			return true
		}
	}
	return false
}
