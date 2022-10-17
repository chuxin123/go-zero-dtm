package logic

import (
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"go-zero-dtm/common"
	"go-zero-dtm/order/rpc/order"
	orderPb "go-zero-dtm/order/rpc/types/pb"
	"go-zero-dtm/product/rpc/product"
	productPb "go-zero-dtm/product/rpc/types/pb"

	"go-zero-dtm/order/api/internal/svc"
	"go-zero-dtm/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTccOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTccOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTccOrderLogic {
	return &CreateTccOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTccOrderLogic) CreateTccOrder(req *types.OrderCreateReq) (resp *types.OrderCreateReply, err error) {
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
	replyBranch1 := &productPb.CheckStockResp{}
	replyBranch2 := &orderPb.CheckOrderResp{}

	if err = dtmgrpc.TccGlobalTransaction(l.svcCtx.Config.DtmServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
		err = tcc.CallBranch(checkStockReq, productTarget+"/product.Product/checkStock", productTarget+"/product.Product/deductStock", productTarget+"/product.Product/deductStockRollback", replyBranch1)
		if err != nil {
			return err
		}
		return tcc.CallBranch(createOrderReq, orderTarget+"/order.Order/checkOrder", orderTarget+"/order.Order/createOrder", orderTarget+"/order.Order/createOrderRollback", replyBranch2)
	}); err != nil {
		return nil, fmt.Errorf("tcc 消息失败:%s", err.Error())
	}

	if replyBranch2.Result && replyBranch1.Result {
		return &types.OrderCreateReply{
			Result:  true,
			OrderNo: createOrderReq.OrderNo,
			Message: "订单正在创建中...",
		}, nil
	}
	return &types.OrderCreateReply{
		Result:  false,
		Message: "订单创建失败...",
	}, nil
}
