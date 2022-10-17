package logic

import (
	"context"
	"go-zero-dtm/product/rpc/internal/svc"
	"go-zero-dtm/product/rpc/types/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckStockLogic {
	return &CheckStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckStockLogic) CheckStock(in *pb.CheckStockReq) (*pb.CheckStockResp, error) {
	// todo: add your logic here and delete this line

	product, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.ProductId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if product.StockNums >= in.Nums {
		return &pb.CheckStockResp{
			Result: true,
		}, nil
	} else {
		return &pb.CheckStockResp{
			Result: false,
		}, nil
	}
}
