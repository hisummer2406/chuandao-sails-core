package sf

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/logic/sf"
	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/callback/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 骑士撤单状态回调
func SfRiderRecallHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SFRiderRecallReq
		if err := httpx.Parse(r, &req); err != nil {
			//参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := sf.NewSfRiderRecallLogic(r.Context(), svcCtx)
		resp, err := l.SfRiderRecall(&req)

		//统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
