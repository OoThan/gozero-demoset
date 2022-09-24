package logic

import (
	"context"
	"go-zero_microservices/mall/service/order/model"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/order/rpc/internal/svc"
	"go-zero_microservices/mall/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *order.DetailRequest) (*order.DetailResponse, error) {
	// order detail
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "order not found!")
		}
		return nil, status.Error(500, err.Error())
	}

	return &order.DetailResponse{
		Id:     res.Id,
		Pid:    res.Pid,
		Uid:    res.Uid,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
}
