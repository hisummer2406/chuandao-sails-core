package utils

import (
	"chuandao-sails-core/apps/platform-gateway/pkg/constants"
	"strconv"
)

func ParseMilliseconds(ms string) int64 {
	if ms == "" {
		return 0
	}
	timestamp, _ := strconv.ParseInt(ms, 10, 64)
	return timestamp / 1000 // 毫秒转秒
}

func ParseTimestamp(ts string) int64 {
	if ts == "" {
		return 0
	}
	timestamp, _ := strconv.ParseInt(ts, 10, 64)
	return timestamp
}

func ParseWeight(weight string) int64 {
	if weight == "" {
		return 0
	}
	w, _ := strconv.ParseFloat(weight, 64)
	return int64(w)
}

func GetSendTypeFromSF(isReverseOrder bool) int {
	if isReverseOrder {
		return constants.SEND_TYPE_HELP_GET // 帮我取
	}
	return constants.SEND_TYPE_HELP_SEND // 帮我送（默认）
}
