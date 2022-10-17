package logic

import (
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"go-zero-dtm/common"
	"go-zero-dtm/order/api/internal/svc"
	"go-zero-dtm/order/api/internal/types"
	"go-zero-dtm/order/rpc/order"
	"go-zero-dtm/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSagaOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSagaOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSagaOrderLogic {
	return &CreateSagaOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSagaOrderLogic) CreateSagaOrder(req *types.OrderCreateReq) (resp *types.OrderCreateReply, err error) {
	orderTarget, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, fmt.Errorf("获取orderTarget失败")
	}
	productTarget, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, fmt.Errorf("获取productTarget失败")
	}

	checkStockReq := &product.CheckStockReq{
		ProductId: req.ProductId,
		Nums:      req.Nums,
	}

	createOrderReq := &order.CreateOrderReq{
		OrderNo:     common.Uniqid("101037"),
		ProductId:   req.ProductId,
		ProductName: req.ProductName,
		UserId:      req.UserId,
		Nums:        req.Nums,
	}
	logx.Infof("createOrderReq: %#v", createOrderReq)

	gid := dtmgrpc.MustGenGid(l.svcCtx.Config.DtmServer)
	saga := dtmgrpc.NewSagaGrpc(l.svcCtx.Config.DtmServer, gid).
		Add(productTarget+"/product.Product/deductStock", productTarget+"/product.Product/deductStockRollback", checkStockReq).
		Add(orderTarget+"/order.Order/createOrder", orderTarget+"/order.Order/createOrderRollback", createOrderReq)

	err = saga.Submit()
	if err != nil {
		return &types.OrderCreateReply{
			Result:  false,
			Message: "订单创建失败...",
		}, nil
	}
	return &types.OrderCreateReply{
		Result:  true,
		OrderNo: createOrderReq.OrderNo,
		Message: "订单正在创建中...",
	}, nil
}
