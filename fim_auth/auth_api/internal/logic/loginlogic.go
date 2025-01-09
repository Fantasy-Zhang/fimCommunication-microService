package logic

import (
	"context"
	"errors"
	"fimCommunication-microService/fim_auth/auth_model"
	"fimCommunication-microService/utils/jwts"
	"fimCommunication-microService/utils/pwd"

	"fimCommunication-microService/fim_auth/auth_api/internal/svc"
	"fimCommunication-microService/fim_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	//
	var user auth_model.UserModel
	//err = l.svcCtx.DB.Where("user_name = ?", req.UserName).First(&user).Error
	err = l.svcCtx.DB.Take(&user, "id = ?", req.UserName).Error
	if err != nil {
		err = errors.New("用户名或密码错误")
		return
	}
	if !pwd.CheckPwd(user.Pwd, req.Password) {
		err = errors.New("用户名或密码错误")
		return
	}
	token, err := jwts.GenToken(jwts.JwtPayload{
		UserId:   user.ID,
		Nickname: user.Nickname,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		logx.Error(err)
		//err = errors.New("服务内部错误")
		return
	}

	return &types.LoginResponse{Token: token}, nil
}
