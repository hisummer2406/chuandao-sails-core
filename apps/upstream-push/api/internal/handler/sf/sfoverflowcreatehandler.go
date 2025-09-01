package sf

import (
	"net/http"

	"chuandao-sails-core/apps/upstream-push/api/internal/logic/sf"
	"chuandao-sails-core/apps/upstream-push/api/internal/svc"
	"chuandao-sails-core/apps/upstream-push/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 订单创建
func SfOverflowCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SfOverflowCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sf.NewSfOverflowCreateLogic(r.Context(), svcCtx)
		resp, err := l.SfOverflowCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
