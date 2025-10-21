package handler

import (
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/types"
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/logic"
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查询运费
func priceQuoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PriceQuotaReq
		if err := httpx.Parse(r, &req); err != nil {
			//参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := logic.NewPriceQuoteLogic(r.Context(), svcCtx)
		resp, err := l.PriceQuote(&req)

		//统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
