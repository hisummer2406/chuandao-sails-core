package sf

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/app/api/callback/internal/logic/sf"
	"chuandao-sails-core/app/api/callback/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 自动建店回调
func SfAutoCreateShopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sf.NewSfAutoCreateShopLogic(r.Context(), svcCtx)
		err := l.SfAutoCreateShop()

		// 统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
