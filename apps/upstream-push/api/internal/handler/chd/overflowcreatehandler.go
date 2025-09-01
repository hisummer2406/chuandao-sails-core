package chd

import (
	"net/http"

	"chuandao-sails-core/apps/upstream-push/api/internal/logic/chd"
	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 订单创建
func OverflowCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OverflowCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chd.NewOverflowCreateLogic(r.Context(), svcCtx)
		resp, err := l.OverflowCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
