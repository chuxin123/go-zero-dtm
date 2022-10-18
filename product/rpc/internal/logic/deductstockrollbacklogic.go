package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-zero-dtm/product/rpc/internal/svc"
	"go-zero-dtm/product/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductStockRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductStockRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductStockRollbackLogic {
	return &DeductStockRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductStockRollbackLogic) DeductStockRollback(in *pb.CheckStockReq) (*pb.CheckStockResp, error) {
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
		product.StockNums = product.StockNums + nums
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
