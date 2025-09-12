package chd

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/upstream-gateway/api/internal/logic/chd"
	"chuandao-sails-core/apps/upstream-gateway/api/internal/svc"
	"chuandao-sails-core/apps/upstream-gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 加小费
func ChdAddOnlineFeeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChdAddOnlineFeeRequest
		if err := httpx.Parse(r, &req); err != nil {
			//参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := chd.NewChdAddOnlineFeeLogic(r.Context(), svcCtx)
		resp, err := l.ChdAddOnlineFee(&req)

		//统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
