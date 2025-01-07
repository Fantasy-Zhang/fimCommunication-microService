package handler

import (
	"fimCommunication-microService/common/response"
	"fimCommunication-microService/fim_auth/auth_api/internal/logic"
	"fimCommunication-microService/fim_auth/auth_api/internal/svc"
	"net/http"
)

func open_loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewOpen_loginLogic(r.Context(), svcCtx)
		resp, err := l.Open_login()
		response.Response(r, w, resp, err)
	}
}
