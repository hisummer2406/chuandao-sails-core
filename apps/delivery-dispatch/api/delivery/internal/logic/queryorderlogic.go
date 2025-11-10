package logic

import (
	"context"

	"chuandao-sails-core/apps/delivery-dispatch/api/delivery/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/delivery/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询订单详情
func NewQueryOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderLogic {
	return &QueryOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOrderLogic) QueryOrder(req *types.QueryOrderReq) (resp *types.QueryOrderResp, err error) {
	//TODO 查询订单详情接口，查询后同步订单库

	//1.查询订单主表
	order , err :=
	//2.查询最新状态日志
	//3.响应结构

	return
}

// getStatusText 获取订单状态文字
func getStatusText(status string) string {
	statusMap := map[string]string{
		"created":    "已创建",
		"dispatched": "已派单",
		"accepted":   "已接单",
		"picked":     "已取货",
		"delivered":  "已送达",
		"completed":  "已完成",
		"cancelled":  "已取消",
		"failed":     "失效",
	}
	if text, ok := statusMap[status]; ok {
		return text
	}
	return "未知状态"
}
