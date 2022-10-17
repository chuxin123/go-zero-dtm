package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dtm/product/rpc/internal/config"
	"go-zero-dtm/product/rpc/model"
)

type ServiceContext struct {
	Config config.Config

	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
