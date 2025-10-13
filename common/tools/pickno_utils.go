package tools

import (
	"fmt"
	"time"
)

// ProcessPickNo 处理取餐号
func ProcessPickNo(pickNo string) string {
	if pickNo != "" {
		return pickNo
	}

	// 生成取餐号: 时间戳后6位
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%06d", timestamp%1000000)
}
