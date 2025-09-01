package handler

import (
	"chuandao-sails-core/common/response"
	"net/http"

	"chuandao-sails-core/demo/restful/user/internal/logic"
	"chuandao-sails-core/demo/restful/user/internal/svc"
	"chuandao-sails-core/demo/restful/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		l := logic.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello(&req)

		//统一处理响应
		response.HandleResponse(w, r, resp, err)
	}
}
