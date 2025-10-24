package dd

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/logic/dd"
	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 订单状态回调，新增订单接口callback传入的URL
func DdCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DDCallbackReq
		if err := httpx.Parse(r, &req); err != nil {
			//参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := dd.NewDdCallbackLogic(r.Context(), svcCtx)
		resp, err := l.DdCallback(&req)

		//统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
