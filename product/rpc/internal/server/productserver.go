// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package server

import (
	"context"

	"go-zero-dtm/product/rpc/internal/logic"
	"go-zero-dtm/product/rpc/internal/svc"
	"go-zero-dtm/product/rpc/types/pb"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) DeductStock(ctx context.Context, in *pb.CheckStockReq) (*pb.CheckStockResp, error) {
	l := logic.NewDeductStockLogic(ctx, s.svcCtx)
	return l.DeductStock(in)
}

func (s *ProductServer) DeductStockRollback(ctx context.Context, in *pb.CheckStockReq) (*pb.CheckStockResp, error) {
	l := logic.NewDeductStockRollbackLogic(ctx, s.svcCtx)
	return l.DeductStockRollback(in)
}

func (s *ProductServer) CheckStock(ctx context.Context, in *pb.CheckStockReq) (*pb.CheckStockResp, error) {
	l := logic.NewCheckStockLogic(ctx, s.svcCtx)
	return l.CheckStock(in)
}
