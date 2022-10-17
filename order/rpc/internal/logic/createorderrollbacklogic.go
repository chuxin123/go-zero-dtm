package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dtm/order/rpc/internal/svc"
	"go-zero-dtm/order/rpc/types/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderRollbackLogic {
	return &CreateOrderRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderRollbackLogic) CreateOrderRollback(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	// todo: add your logic here and delete this line
	order, err := l.svcCtx.OrderModel.FindOneByOrderNo(l.ctx, in.OrderNo)
	if err == sqlx.ErrNotFound {
		return &pb.CreateOrderResp{}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		order.Delete = 1
		err = l.svcCtx.OrderModel.Update(l.ctx, order)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateOrderResp{
		Id: order.Id,
	}, nil
}
