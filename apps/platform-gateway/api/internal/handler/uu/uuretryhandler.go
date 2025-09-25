package uu

import (
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/platform-gateway/api/internal/logic/uu"
	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 订单回推
func UuRetryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UURetryRequest
		if err := httpx.Parse(r, &req); err != nil {
			//参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := uu.NewUuRetryLogic(r.Context(), svcCtx)
		resp, err := l.UuRetry(&req)

		//统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
