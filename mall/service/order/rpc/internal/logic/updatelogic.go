package logic

import (
	"context"
	"go-zero_microservices/mall/service/order/model"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/order/rpc/internal/svc"
	"go-zero_microservices/mall/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *order.UpdateRequest) (*order.UpdateResponse, error) {
	// order update
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "order not found")
		}
		return nil, status.Error(500, err.Error())
	}

	if res.Uid != 0 {
		res.Uid = in.Uid
	}
	if res.Pid != 0 {
		res.Pid = in.Pid
	}
	if res.Amount != 0 {
		res.Amount = in.Amount
	}
	if res.Status != 0 {
		res.Status = in.Status
	}

	err = l.svcCtx.OrderModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.UpdateResponse{}, nil
}
