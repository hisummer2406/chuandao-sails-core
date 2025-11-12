package ss

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/app/api/callback/internal/logic/ss"
	"chuandao-sails-core/app/api/callback/internal/svc"
	"chuandao-sails-core/app/api/callback/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 订单状态回调，回调账号配置的地址（账户中心->应用中心->回调URL：NotifyUrl）来通知订单状态的变化。
func DdCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SSCallbackReq
		if err := httpx.Parse(r, &req); err != nil {
			// 参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := ss.NewDdCallbackLogic(r.Context(), svcCtx)
		resp, err := l.DdCallback(&req)

		// 统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
