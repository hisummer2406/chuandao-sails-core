package sf

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/platform-gateway/api/internal/logic/sf"
	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"chuandao-sails-core/apps/platform-gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 询价
func SfGetPriceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SfGetPriceRequest
		if err := httpx.Parse(r, &req); err != nil {
			//参数错误
			response.SFHandlerResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := sf.NewSfGetPriceLogic(r.Context(), svcCtx)
		resp, err := l.SfGetPrice(&req)

		//统一处理响应
		response.SFHandlerResponse(w, r, resp, err)
	}
}
