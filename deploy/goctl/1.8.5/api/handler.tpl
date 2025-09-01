package {{.PkgName}}

import (
    "chuandao-sails-core/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
		    //参数错误
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), err.Error()))
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})

		//统一处理响应
        response.HandleResponse(w, r, resp, err)
	}
}
