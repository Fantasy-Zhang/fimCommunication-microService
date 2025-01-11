package logic

import (
	"context"

	"fimCommunication-microService/fim_user/user_rpc/internal/svc"
	"fimCommunication-microService/fim_user/user_rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &user_rpc.UserCreateResponse{}, nil
}
