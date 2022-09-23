package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero_microservices/mall/service/order/model"
	"go-zero_microservices/mall/service/order/rpc/internal/config"
	"go-zero_microservices/mall/service/product/rpc/productclient"
	"go-zero_microservices/mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel
	UserRpc    userclient.User
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
