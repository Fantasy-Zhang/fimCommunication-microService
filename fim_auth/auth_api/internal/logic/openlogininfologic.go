package logic

import (
	"context"

	"fimCommunication-microService/fim_auth/auth_api/internal/svc"
	"fimCommunication-microService/fim_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Open_login_infoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpen_login_infoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Open_login_infoLogic {
	return &Open_login_infoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Open_login_infoLogic) Open_login_info() (resp []types.OpenLoginInfoResponse, err error) {
	// todo: add your logic here and delete this line
	for _, i := range l.svcCtx.Config.OpenLoginList {
		resp = append(resp, types.OpenLoginInfoResponse{
			Href: i.Href,
			Icon: i.Icon,
			Name: i.Name,
		})
	}

	return
}
