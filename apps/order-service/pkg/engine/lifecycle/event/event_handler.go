package event

import (
	"context"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

// Handler 事件处理接口
type Handler interface {
	Handle(ctx context.Context, msg *OrderEventMessage) error
	EventTypes() []string
}

// HandlerManager 事件处理管理器
type HandlerManager struct {
	handlers map[string][]Handler
	logger   logx.Logger
}

func NewHandlerManager() *HandlerManager {
	return &HandlerManager{
		handlers: make(map[string][]Handler),
		logger:   logx.WithContext(context.Background()),
	}
}

// Register 注册处理器
func (h *HandlerManager) Register(handler Handler) {
	for _, eventType := range handler.EventTypes() {
		if _, exists := h.handlers[eventType]; !exists {
			h.handlers[eventType] = []Handler{} // 初始化为一个空的 Handler 切片
		}
		h.handlers[eventType] = append(h.handlers[eventType], handler) // 正确地将 handler 追加到切片
	}
}

func (h *HandlerManager) HandleMessage(ctx context.Context, msgView *rmq_client.MessageView) error {
	//解析事件消息
	var eventMsg OrderEventMessage
	if err := jsonx.Unmarshal(msgView.GetBody(), &eventMsg); err != nil {
		h.logger.Errorf("Failed to unmarshal order events message: %v", err)
		return err
	}

	//获取事件处理器
	handlers, exists := h.handlers[eventMsg.EventType]
	if !exists {
		h.logger.Errorf("Unknown events type: %s", eventMsg.EventType)
		return nil
	}

	//依次执行处理器
	for _, handler := range handlers {
		if err := handler.Handle(ctx, &eventMsg); err != nil {
			h.logger.Errorf("Failed to handle events: %v", err)
			return err
		}
	}
	return nil
}
