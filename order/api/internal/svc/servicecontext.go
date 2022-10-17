package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dtm/order/api/internal/config"
	"go-zero-dtm/order/rpc/order"
	"go-zero-dtm/product/rpc/product"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc   order.Order
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
