package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dtm/product/rpc/internal/svc"
	"go-zero-dtm/product/rpc/types/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductStockLogic {
	return &DeductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductStockLogic) DeductStock(in *pb.CheckStockReq) (*pb.CheckStockResp, error) {
	// todo: add your logic here and delete this line
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	productId := in.ProductId
	nums := in.Nums

	product, err := l.svcCtx.ProductModel.FindOne(l.ctx, productId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		product.StockNums = product.StockNums - nums
		err = l.svcCtx.ProductModel.Update(l.ctx, product)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CheckStockResp{
		Result: true,
	}, nil
}
