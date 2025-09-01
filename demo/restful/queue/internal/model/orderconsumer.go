package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

func (c *Consumer) handleOrderCreated(ctx context.Context, order *OrderMessage) error {
	logx.Infof("rocketmq processing order create: %s", order.OrderID)
	// TODO: 处理订单创建业务逻辑
	// - 保存订单到数据库
	// - 发送库存锁定消息
	// - 发送用户通知
	return nil
}
