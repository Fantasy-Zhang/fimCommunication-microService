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

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) (resp string, err error) {
	//
	if token == "" {
		err = errors.New("请传入token")
		return
	}
	payload, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("token错误")
		return
	}
	now := time.Now()
	expiration := payload.ExpiresAt.Time.Sub(now)
	key := fmt.Sprintf("logout_%d", payload.UserId)
	//过期时间为jwt失效时间
	l.svcCtx.Redis.SetNX(l.ctx, key, "", expiration)
	resp = "注销成功"
	return
}
