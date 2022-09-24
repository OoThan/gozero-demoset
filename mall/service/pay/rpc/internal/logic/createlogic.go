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

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pay.CreateRequest) (*pay.CreateResponse, error) {
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
	_, err = l.svcCtx.PayModel.FindOneByOid(l.ctx, in.Oid)
	if err != nil {
		return nil, status.Error(500, "Order has already paid")
	}

	newPay := model.Pay{
		Uid:    in.Uid,
		Oid:    in.Oid,
		Amount: in.Amount,
		Source: 0,
		Status: 0,
	}

	res, err := l.svcCtx.PayModel.Insert(l.ctx, &newPay)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pay.CreateResponse{
		Id: uint64(insertedId),
	}, nil
}
