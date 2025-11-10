package handler

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/delivery-dispatch/api/delivery/internal/logic"
	"chuandao-sails-core/apps/delivery-dispatch/api/delivery/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func healthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHealthLogic(r.Context(), svcCtx)
		resp, err := l.Health()

		// 统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
