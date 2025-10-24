package sf

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/logic/sf"
	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 订单完成回调
func SfOrderCompleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SFOrderCompleteReq
		if err := httpx.Parse(r, &req); err != nil {
			//参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := sf.NewSfOrderCompleteLogic(r.Context(), svcCtx)
		resp, err := l.SfOrderComplete(&req)

		//统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
