package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	OrderRpc   zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf

	DtmServer string
}
