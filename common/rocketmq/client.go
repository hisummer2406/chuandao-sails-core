package rocketmq

import (
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

type Client struct {
	config   *Config
	producer rmq_client.Producer
	consumer rmq_client.SimpleConsumer
	handlers map[string]MessageHandler
	mu       sync.Mutex
	closed   bool
}

// NewClient 创建新的RocketMQ客户端
func NewClient(config *Config) (*Client, error) {
	if config == nil {
		return nil, fmt.Errorf("rocketmq config is nil")
	}

	client := &Client{
		config:   config,
		handlers: make(map[string]MessageHandler),
	}

	//初始化生产者
	if err := client.initProducer(); err != nil {
		return nil, fmt.Errorf("rocketmq init producer failed: %v", err)
	}

	if err := client.initConsumer(); err != nil {
		return nil, fmt.Errorf("rocketmq init consumer failed: %v", err)
	}

	return client, nil
}

// initProducer 初始化生产者
func (c *Client) initProducer() error {
	creds := &credentials.SessionCredentials{
		AccessKey:    c.config.AccessKey,
		AccessSecret: c.config.SecretKey,
	}

	producer, err := rmq_client.NewProducer(&rmq_client.Config{
		Endpoint:    c.config.Endpoint,
		NameSpace:   c.config.NameSpace,
		Credentials: creds,
	}, rmq_client.WithTopics(c.config.Producer.Topics...))

	if err != nil {
		return err
	}

	if err := producer.Start(); err != nil {
		return err
	}

	c.producer = producer
	return nil
}

// initConsumer 初始化消费者
func (c *Client) initConsumer() error {
	creds := &credentials.SessionCredentials{
		AccessKey:    c.config.AccessKey,
		AccessSecret: c.config.SecretKey,
	}

	consumer, err := rmq_client.NewSimpleConsumer(&rmq_client.Config{
		Endpoint:    c.config.Endpoint,
		NameSpace:   c.config.NameSpace,
		Credentials: creds,
	}, rmq_client.WithAwaitDuration(c.config.Consumer.AwaitDuration))

	if err != nil {
		return err
	}

	c.consumer = consumer
	return nil
}

// Send 发送消息
func (c *Client) Send(ctx context.Context, msg BusinessMessage) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return fmt.Errorf("rocketmq client is closed")
	}

	body, err := msg.ToBytes()
	if err != nil {
		return fmt.Errorf("rocketmq message to bytes failed: %v", err)
	}

	message := &rmq_client.Message{
		Topic: msg.GetTopic(),
		Body:  body,
	}

	//设置消息属性
	if tag := msg.GetTag(); tag != "" {
		message.SetTag(tag)
	}

	//设置消息键
	if keys := msg.GetKeys(); len(keys) > 0 {
		message.SetKeys(keys...)
	}

	_, err = c.producer.Send(ctx, message)
	return err
}

// SendWithOptions 发送消息
func (c *Client) sendWithOptions(ctx context.Context, topic string, body []byte, opts ...MessageOption) error {

	message := &rmq_client.Message{
		Topic: topic,
		Body:  body,
	}

	for _, opt := range opts {
		opt(message)
	}

	_, err := c.producer.Send(ctx, message)
	return err
}

// SendAsync 异步发送
func (c *Client) SendAsync(ctx context.Context, msg BusinessMessage, callback func(error)) {
	go func() {
		err := c.Send(ctx, msg)
		if callback != nil {
			callback(err)
		}
	}()
}

func (c *Client) Subscribe(topic string, handler MessageHandler) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return fmt.Errorf("rocketmq client is closed")
	}

	filterExpr := rmq_client.SUB_ALL
	if err := c.consumer.Subscribe(topic, filterExpr); err != nil {
		return fmt.Errorf("rocketmq consumer subscribe failed: %v", err)
	}

	c.handlers[topic] = handler
	return nil
}

// StartConsumer 启动消费者
func (c *Client) StartConsumer() error {
	if err := c.consumer.Start(); err != nil {
		return fmt.Errorf("rocketmq consumer start failed: %v", err)
	}

	go c.consumerLoop()
	return nil

}

// consumerLoop 消费者循环
func (c *Client) consumerLoop() {
	for !c.closed {
		//接收消息
		messages, err := c.consumer.Receive(
			context.Background(),
			c.config.Consumer.MaxMessageNum,
			c.config.Consumer.InvisibleDuration,
		)
		if err != nil {
			logx.Errorf("rocketmq consumer receive failed: %v", err)
			time.Sleep(time.Second)
			continue
		}

		//处理消息
		for _, msg := range messages {
			c.handleMessage(msg)
		}
	}
}

// handleMessage 处理单条消息
func (c *Client) handleMessage(msg *rmq_client.MessageView) {
	topic := msg.GetTopic()

	c.mu.Lock()
	handler, exists := c.handlers[topic]
	c.mu.Unlock()

	if !exists {
		logx.Errorf("rocketmq consumer topic not found: %s", topic)
		return
	}

	//创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := handler(ctx, msg); err != nil {
		logx.Errorf("rocketmq consumer handler failed: %v", err)
		return
	}

	//确认消息
	if err := c.consumer.Ack(context.Background(), msg); err != nil {
		logx.Errorf("rocketmq consumer ack failed: %v", err)
		return
	}
}

// Close 关闭客户端
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil
	}

	c.closed = true

	var errs []error

	if c.producer != nil {
		if err := c.producer.GracefulStop(); err != nil {
			errs = append(errs, fmt.Errorf("rocketmq producer graceful stop failed: %v", err))
		}
	}

	if c.consumer != nil {
		if err := c.consumer.GracefulStop(); err != nil {
			errs = append(errs, fmt.Errorf("rocketmq consumer graceful stop failed: %v", err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("rocketmq consumer failed: %v", errs)
	}

	return nil
}
