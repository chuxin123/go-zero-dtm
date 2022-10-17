package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"go-zero-dtm/order/rpc/internal/svc"
	"go-zero-dtm/order/rpc/types/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckOrderLogic {
	return &CheckOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckOrderLogic) CheckOrder(in *pb.CreateOrderReq) (*pb.CheckOrderResp, error) {

	orderNo := in.OrderNo
	_, err := l.svcCtx.OrderModel.FindOneByOrderNo(l.ctx, orderNo)
	if err == sqlc.ErrNotFound {
		return &pb.CheckOrderResp{
			Result: true,
		}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CheckOrderResp{
		Result: false,
	}, nil
}
