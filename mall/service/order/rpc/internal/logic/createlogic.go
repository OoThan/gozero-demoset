package logic

import (
	"context"
	"go-zero_microservices/mall/service/order/model"
	"go-zero_microservices/mall/service/product/rpc/product"
	"go-zero_microservices/mall/service/user/rpc/user"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/order/rpc/internal/svc"
	"go-zero_microservices/mall/service/order/rpc/order"

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

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	// check user
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: int64(in.Uid),
	})
	if err != nil {
		return nil, err
	}

	// check product
	productRes, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{
		Id: in.Pid,
	})
	if err != nil {
		return nil, err
	}
	if productRes.Stock <= 0 {
		return nil, status.Error(500, "Insufficient product stock")
	}
	newOrder := model.Order{
		Uid:    in.Uid,
		Pid:    in.Pid,
		Amount: in.Amount,
		Status: 0,
	}

	// create order
	res, err := l.svcCtx.OrderModel.Insert(l.ctx, &newOrder)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Inserted Order ID
	insertedId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// update product detail
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
		Id:     productRes.Id,
		Name:   productRes.Name,
		Desc:   productRes.Desc,
		Stock:  productRes.Stock - 1,
		Amount: productRes.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &order.CreateResponse{
		Id: uint64(insertedId),
	}, nil
}
