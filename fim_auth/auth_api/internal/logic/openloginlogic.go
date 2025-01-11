package logic

import (
	"context"
	"errors"
	"fimCommunication-microService/fim_auth/auth_model"
	user_rpc "fimCommunication-microService/fim_user/user_rpc/types/user"
	"fimCommunication-microService/utils/open_login"
	"fmt"

	"fimCommunication-microService/fim_auth/auth_api/internal/svc"
	"fimCommunication-microService/fim_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Open_loginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpen_loginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Open_loginLogic {
	return &Open_loginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Open_loginLogic) Open_login(req *types.OpenLoginRequest) (resp *types.LoginResponse, err error) {
	switch req.Flag {
	case "qq":
		info, err := open_login.NewQQLogin(req.Code, open_login.QQConfig{
			AppID:    l.svcCtx.Config.QQ.AppID,
			AppKey:   l.svcCtx.Config.QQ.AppKey,
			Redirect: l.svcCtx.Config.QQ.Redirect,
		})
		if err != nil {
			logx.Error(err)
			return nil, errors.New("qq登录失败")
		}
		var user auth_model.UserModel
		if err := l.svcCtx.DB.Where("open_id = ?", info.OpenID).First(&user).Error; err != nil {
			//logx.Error(err)
			//return nil, errors.New("用户不存在")
			fmt.Println("注册失败")
			res, err := l.svcCtx.UserRpc.UserCreate(context.Background(), &user_rpc.UserCreateRequest{
				NickName: info.Nickname,
				Password: "",
				Role:     2,
				Avatar:   info.Avatar,
				OpenId:   info.OpenID,
			})
			if err != nil {
				logx.Error(err)
				return nil, errors.New("注册失败")
			}
			user.Model.ID = uint(res.UserId)
			user.Nickname = info.Nickname
			user.Role = 2
		}
		//登录逻辑
		//token, err := jwts.GenToken(jwts.JwtPayload{
		//	UserId:   user.ID,
		//	Nickname: user.Nickname,
		//	Role:     user.Role,
		//}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	case "wx":
		//微信登录
	case "wb":
		//微博登录
	}

	return
}
