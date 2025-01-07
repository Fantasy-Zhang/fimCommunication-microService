package handler

import (
	"fimCommunication-microService/common/response"
	"fimCommunication-microService/fim_auth/auth_api/internal/logic"
	"fimCommunication-microService/fim_auth/auth_api/internal/svc"
	"net/http"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication()
		response.Response(r, w, resp, err)
	}
}
