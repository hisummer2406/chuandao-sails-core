package uu

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/app/api/callback/internal/logic/uu"
	"chuandao-sails-core/app/api/callback/internal/svc"
	"chuandao-sails-core/app/api/callback/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 订单状态回调，用户提交订单时传入
func UuCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUCallbackReq
		if err := httpx.Parse(r, &req); err != nil {
			// 参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := uu.NewUuCallbackLogic(r.Context(), svcCtx)
		resp, err := l.UuCallback(&req)

		// 统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
