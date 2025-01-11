package svc

import (
	"fimCommunication-microService/core"
	"fimCommunication-microService/fim_auth/auth_api/internal/config"
	"fimCommunication-microService/fim_user/user_rpc/types/user"
	"fimCommunication-microService/fim_user/user_rpc/users"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc user_rpc.UsersClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	client := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	return &ServiceContext{
		Config:  c,
		DB:      mysqlDb,
		Redis:   client,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
	}
}
