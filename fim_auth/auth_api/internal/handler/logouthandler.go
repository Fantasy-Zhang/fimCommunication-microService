package handler

import (
	"net/http"

	"fimCommunication-microService/fim_auth/auth_api/internal/logic"
	"fimCommunication-microService/fim_auth/auth_api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}