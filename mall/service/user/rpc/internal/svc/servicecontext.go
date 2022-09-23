package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero_microservices/mall/service/user/model"
	"go-zero_microservices/mall/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
