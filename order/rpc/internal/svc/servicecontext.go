package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dtm/order/rpc/internal/config"
	"go-zero-dtm/order/rpc/model"
)

type ServiceContext struct {
	Config config.Config

	OrderModel model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
