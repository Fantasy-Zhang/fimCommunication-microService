package logic

import (
	"context"
	"errors"
	"fimCommunication-microService/utils/jwts"
	"fmt"
	"time"

	"fimCommunication-microService/fim_auth/auth_api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(token string) (resp string, err error) {
	if token == "" {
		err = errors.New("认证失败")
		return
	}
	payload, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("认证失败")
		return
	}
	l.svcCtx.Redis.SetNX(l.ctx, fmt.Sprintf("logout_%d", payload.UserId), "", time.Duration(payload.ExpiresAt.Unix())*time.Second)
	_, err = l.svcCtx.Redis.Get(l.ctx, fmt.Sprintf("logout_%d", payload.UserId)).Result()
	if err != nil {
		err = errors.New("认证失败")
		return
	}
	resp = "认证通过"
	return resp, nil
}
