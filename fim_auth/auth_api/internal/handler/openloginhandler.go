package handler

import (
	"fimCommunication-microService/common/response"
	"fimCommunication-microService/fim_auth/auth_api/internal/logic"
	"fimCommunication-microService/fim_auth/auth_api/internal/svc"
	"fimCommunication-microService/fim_auth/auth_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func open_loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OpenLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOpen_loginLogic(r.Context(), svcCtx)
		resp, err := l.Open_login(&req)
		response.Response(r, w, resp, err)
	}
}
