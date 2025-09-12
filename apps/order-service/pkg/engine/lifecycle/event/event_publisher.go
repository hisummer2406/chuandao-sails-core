package event

import "github.com/zeromicro/go-zero/core/logx"

type EventPublisher struct {
	producer rocketmq.Producer
	topic    string
	logger   logx.Logger
}
