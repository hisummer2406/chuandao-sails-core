package sf

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/platform-gateway/api/internal/logic/sf"
	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"
	"chuandao-sails-core/apps/platform-gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 订单创建
func SfCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SfCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			//参数错误
			response.SFHandlerResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := sf.NewSfCreateLogic(r.Context(), svcCtx)
		resp, err := l.SfCreate(&req)

		//统一处理响应
		response.SFHandlerResponse(w, r, resp, err)
	}
}
