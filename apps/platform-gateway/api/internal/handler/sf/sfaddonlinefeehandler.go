package sf

import (
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/platform-gateway/api/internal/logic/sf"
	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 加小费
func SfAddOnlineFeeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SfAddOnlineFeeRequest
		if err := httpx.Parse(r, &req); err != nil {
			//参数错误
			response.SFHandlerResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := sf.NewSfAddOnlineFeeLogic(r.Context(), svcCtx)
		resp, err := l.SfAddOnlineFee(&req)

		//统一处理响应
		response.SFHandlerResponse(w, r, resp, err)
	}
}
