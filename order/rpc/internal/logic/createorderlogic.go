package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dtm/order/rpc/internal/svc"
	"go-zero-dtm/order/rpc/model"
	"go-zero-dtm/order/rpc/types/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	// todo: add your logic here and delete this line

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	orderCreate := &pb.CreateOrderResp{}
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		orderReq := &model.Order{
			OrderNo:     in.OrderNo,
			UserId:      in.UserId,
			ProductId:   in.ProductId,
			ProductName: in.ProductName,
			Nums:        in.Nums,
			Delete:      0,
		}
		result, err := l.svcCtx.OrderModel.Insert(l.ctx, orderReq)
		if err != nil {
			return fmt.Errorf("创建订单失败 err : %v , order:%+v \n", err, orderReq)
		}

		lastId, err := result.LastInsertId()
		orderCreate.Id = lastId
		if err != nil {
			return fmt.Errorf("创建订单失败 err : %v , order:%+v \n", err, orderReq)
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateOrderResp{
		Id: orderCreate.Id,
	}, nil
}
