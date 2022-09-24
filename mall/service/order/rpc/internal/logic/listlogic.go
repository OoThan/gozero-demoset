package logic

import (
	"context"
	"go-zero_microservices/mall/service/order/model"
	"go-zero_microservices/mall/service/user/rpc/user"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/order/rpc/internal/svc"
	"go-zero_microservices/mall/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *order.ListRequest) (*order.ListResponse, error) {
	// check user
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: int64(in.Uid),
	})
	if err != nil {
		return nil, err
	}

	// list order
	list, err := l.svcCtx.OrderModel.FindAllByUid(l.ctx, in.Uid)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "order not found")
		}
		return nil, status.Error(500, err.Error())
	}

	orderList := make([]*order.DetailResponse, 0)
	for _, item := range list {
		orderList = append(orderList, &order.DetailResponse{
			Id:     item.Id,
			Uid:    item.Uid,
			Pid:    item.Pid,
			Amount: item.Amount,
			Status: item.Status,
		})
	}

	return &order.ListResponse{
		Data: orderList,
	}, nil
}
