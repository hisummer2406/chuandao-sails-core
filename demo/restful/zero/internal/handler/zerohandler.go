package handler

import (
	"net/http"

	"chuandao-sails-core/demo/restful/zero/internal/logic"
	"chuandao-sails-core/demo/restful/zero/internal/svc"
	"chuandao-sails-core/demo/restful/zero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ZeroHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewZeroLogic(r.Context(), svcCtx)
		resp, err := l.Zero(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
