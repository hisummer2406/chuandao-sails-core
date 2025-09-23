package state

import (
	"chuandao-sails-core/apps/order-service/pkg/engine/lifecycle/event"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type StateMachine struct {
	publisher *event.Publisher
	validator *StateValidator
	logger    logx.Logger
}

func NewStateMachine(publisher *event.Publisher) *StateMachine {
	return &StateMachine{
		publisher: publisher,
		validator: NewStateValidator(),
		logger:    logx.WithContext(context.Background()),
	}
}

// StateTransitionRequest 状态转换请求结构
type StateTransitionRequest struct {
	OrderID       string                 `json:"order_id"`       // 订单ID
	CurrentStatus string                 `json:"current_status"` // 当前状态
	TargetStatus  string                 `json:"target_status"`  // 目标状态
	Reason        string                 `json:"reason"`         // 转换原因
	OperatorID    string                 `json:"operator_id"`    // 操作者ID
	OperatorType  string                 `json:"operator_type"`  // 操作者类型：SYSTEM/USER/PLATFORM
	Extra         map[string]interface{} `json:"extra"`          // 额外数据
	Timestamp     time.Time              `json:"timestamp"`      // 操作时间
}

// TransitionTo 状态转换
func (m *StateMachine) TransitionTo(ctx context.Context, req *StateTransitionRequest) error {
	// 1.验证状态转换
	m.validator.IsValidTransition(req.CurrentStatus, req.TargetStatus)
	// 2.更新数据库状态

	// 3.发布状态变更事件

	// 4.发布生命周期事件

	return nil
}
