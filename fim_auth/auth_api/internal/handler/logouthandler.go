package handler

import (
	"fimCommunication-microService/common/response"
	"fimCommunication-microService/fim_auth/auth_api/internal/logic"
	"fimCommunication-microService/fim_auth/auth_api/internal/svc"
	"net/http"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		token := r.Header.Get("token")
		resp, err := l.Logout(token)
		response.Response(r, w, resp, err)
	}
}
