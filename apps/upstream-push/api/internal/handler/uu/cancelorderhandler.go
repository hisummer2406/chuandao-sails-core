package uu

import (
	"net/http"

	"chuandao-sails-core/apps/upstream-push/api/internal/logic/uu"
	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 接收取消订单消息
func CancelOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CancelOrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := uu.NewCancelOrderLogic(r.Context(), svcCtx)
		resp, err := l.CancelOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
