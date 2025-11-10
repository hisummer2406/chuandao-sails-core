package handler

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/delivery-dispatch/api/delivery/internal/logic"
	"chuandao-sails-core/apps/delivery-dispatch/api/delivery/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/delivery/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查询运费
func getQuoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetQuotaReq
		if err := httpx.Parse(r, &req); err != nil {
			// 参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := logic.NewGetQuoteLogic(r.Context(), svcCtx)
		resp, err := l.GetQuote(&req)

		// 统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
