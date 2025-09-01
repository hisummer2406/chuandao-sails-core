package logic

import (
	"chuandao-sails-core/demo/restful/queue/internal/model"
	"chuandao-sails-core/demo/restful/user/internal/svc"
	"chuandao-sails-core/demo/restful/user/internal/types"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.Request) (resp *types.Response, err error) {
	orderMsg := &mq.OrderMessage{
		OrderID:     uuid.New().String(),
		UserID:      req.Name,
		Amount:      100,
		Status:      "created",
		ProductInfo: map[string]interface{}{"name": "test"},
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	if err := l.svcCtx.RocketMQ.Send(l.ctx, orderMsg); err != nil {
		l.Logger.Errorf("rocketmq send order message error: %v", err)
		return nil, fmt.Errorf("rocketmq send order message error: %v", err)
	}

	l.Logger.Infof("rocketmq send order message success, orderID=%s", orderMsg.OrderID)

	return &types.Response{
		Hello: "hello " + req.Name,
	}, nil
}
