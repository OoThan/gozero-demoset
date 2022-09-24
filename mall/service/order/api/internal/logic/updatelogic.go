package logic

import (
	"context"
	"go-zero_microservices/mall/service/order/rpc/orderclient"

	"go-zero_microservices/mall/service/order/api/internal/svc"
	"go-zero_microservices/mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequset) (resp *types.UpdateResponse, err error) {
	// order update
	_, err = l.svcCtx.OrderRpc.Update(l.ctx, &orderclient.UpdateRequest{
		Id:     req.Id,
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{}, nil
}
