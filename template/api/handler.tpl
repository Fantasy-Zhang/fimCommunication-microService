package handler

import (
	"net/http"
    "fimCommunication-microService/common/response"
    {{.ImportPackages}}
    {{if .HasRequest}}
	"github.com/zeromicro/go-zero/rest/httpx"
	{{end}}
)


func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		{{end}}
		l := logic.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
	    {{if .HasResp}}response.Response(r,w, resp, err){{else}}response.Response(w, nil, err){{end}}
	}
}
