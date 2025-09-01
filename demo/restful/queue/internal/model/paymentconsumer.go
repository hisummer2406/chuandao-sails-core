package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

func (c *Consumer) handlePaymentSuccess(ctx context.Context, payment *PaymentMessage) error {
	logx.Infof("rocketmq processing payment success: %s", payment.PaymentID)
	// TODO: 处理支付成功业务逻辑
	// - 更新订单状态
	// - 发送支付成功通知
	return nil
}
