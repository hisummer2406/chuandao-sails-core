package sf

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/app/api/callback/internal/logic/sf"
	"chuandao-sails-core/app/api/callback/internal/svc"
	"chuandao-sails-core/app/api/callback/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 顺丰原因订单取消回调
func SfCancelCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SFCancelCallbackReq
		if err := httpx.Parse(r, &req); err != nil {
			// 参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := sf.NewSfCancelCallbackLogic(r.Context(), svcCtx)
		resp, err := l.SfCancelCallback(&req)

		// 统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
