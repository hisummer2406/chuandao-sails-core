package rocketmq

import (
	"context"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
)

// MessageHandler 消息处理函数
type MessageHandler func(ctx context.Context, msg *rmq_client.MessageView) error

// BusinessMessage 业务消息接口
type BusinessMessage interface {
	GetTopic() string
	GetTag() string
	GetKeys() []string
	ToBytes() ([]byte, error)
}
