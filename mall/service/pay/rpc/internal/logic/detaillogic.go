package logic

import (
	"context"
	"go-zero_microservices/mall/service/pay/model"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/pay/rpc/internal/svc"
	"go-zero_microservices/mall/service/pay/rpc/pay"

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

func (l *DetailLogic) Detail(in *pay.DetailRequest) (*pay.DetailResponse, error) {
	// check payment
	res, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "payment not found")
		}
		return nil, status.Error(500, err.Error())
	}

	return &pay.DetailResponse{
		Id:     res.Id,
		Uid:    res.Uid,
		Oid:    res.Oid,
		Amount: res.Amount,
		Source: res.Source,
		Status: res.Status,
	}, nil
}
