package rocketmq

import (
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"time"
)

// MessageOption 消息选项函数
type MessageOption func(*rmq_client.Message)

// WithTag 设置消息标签
func WithTag(tag string) MessageOption {
	return func(msg *rmq_client.Message) {
		msg.SetTag(tag)
	}
}

// WithKeys 设置消息键值
func WithKeys(keys ...string) MessageOption {
	return func(msg *rmq_client.Message) {
		msg.SetKeys(keys...)
	}
}

// WithDelay 设置消息延迟时间
func WithDelay(delay time.Duration) MessageOption {
	return func(msg *rmq_client.Message) {
		msg.SetDelayTimestamp(time.Now().Add(delay))
	}
}

// WithFIFO 设置消息为FIFO队列
func WithFIFO(group string) MessageOption {
	return func(msg *rmq_client.Message) {
		msg.SetMessageGroup(group)
	}
}

// WithProperty 设置消息属性
func WithProperty(key, value string) MessageOption {
	return func(msg *rmq_client.Message) {
		msg.AddProperty(key, value)
	}
}
