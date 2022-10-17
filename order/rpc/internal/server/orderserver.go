// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package server

import (
	"context"

	"go-zero-dtm/order/rpc/internal/logic"
	"go-zero-dtm/order/rpc/internal/svc"
	"go-zero-dtm/order/rpc/types/pb"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) CheckOrder(ctx context.Context, in *pb.CreateOrderReq) (*pb.CheckOrderResp, error) {
	l := logic.NewCheckOrderLogic(ctx, s.svcCtx)
	return l.CheckOrder(in)
}

func (s *OrderServer) CreateOrder(ctx context.Context, in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	l := logic.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}

func (s *OrderServer) CreateOrderRollback(ctx context.Context, in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	l := logic.NewCreateOrderRollbackLogic(ctx, s.svcCtx)
	return l.CreateOrderRollback(in)
}
