package logic

import (
	"chuandao-sails-core/demo/restful/queue/internal/model"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"chuandao-sails-core/demo/restful/user/internal/svc"
	"chuandao-sails-core/demo/restful/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req *types.HelloRequest) (resp *types.Response, err error) {

	//发送支付成功消息
	paymentMsg := &mq.PaymentMessage{
		PaymentID:   uuid.New().String(),
		OrderID:     req.OrderID,
		UserID:      req.UserID,
		Amount:      req.Amount,
		Method:      "alipay",
		Status:      "created",
		ProcessTime: time.Now(),
	}

	if err := l.svcCtx.RocketMQ.Send(l.ctx, paymentMsg); err != nil {
		l.Logger.Errorf("rocketmq send payment message error: %v", err)
		return nil, fmt.Errorf("rocketmq send payment message error: %v", err)
	}

	//发送订单支付成功消息
	orderMsg := &mq.OrderMessage{
		OrderID:    req.OrderID,
		Status:     "paid",
		UpdateTime: time.Now(),
	}

	if err := l.svcCtx.RocketMQ.Send(l.ctx, orderMsg); err != nil {
		l.Logger.Errorf("rocketmq send order message error: %v", err)
		return nil, fmt.Errorf("rocketmq send order message error: %v", err)
	}

	return &types.Response{
		Hello: "hello " + req.Name,
	}, nil
}
