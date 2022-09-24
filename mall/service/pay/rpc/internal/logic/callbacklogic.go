package logic

import (
	"context"
	"go-zero_microservices/mall/service/order/rpc/order"
	"go-zero_microservices/mall/service/pay/model"
	"go-zero_microservices/mall/service/user/rpc/user"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/pay/rpc/internal/svc"
	"go-zero_microservices/mall/service/pay/rpc/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	// check user
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: int64(in.Uid),
	})
	if err != nil {
		return nil, err
	}

	// check order
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, err
	}

	// check payment
	res, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "payment not found")
		}
		return nil, status.Error(500, err.Error())
	}
	// check payment and order amount
	if in.Amount != res.Amount {
		return nil, status.Error(100, "payment doesn't match order amount")
	}

	res.Source = in.Source
	res.Status = in.Status

	err = l.svcCtx.PayModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// update order payment status
	_, err = l.svcCtx.OrderRpc.Paid(l.ctx, &order.PaidRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pay.CallbackResponse{}, nil
}
