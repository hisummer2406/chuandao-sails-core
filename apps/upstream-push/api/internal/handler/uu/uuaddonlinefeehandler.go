package uu

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/upstream-push/api/internal/logic/uu"
	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 加小费
func UuAddOnlineFeeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUAddOnlineFeeRequest
		if err := httpx.Parse(r, &req); err != nil {
			//参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := uu.NewUuAddOnlineFeeLogic(r.Context(), svcCtx)
		resp, err := l.UuAddOnlineFee(&req)

		//统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
