package logic

import (
	"context"
	"go-zero_microservices/mall/service/product/model"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/product/rpc/internal/svc"
	"go-zero_microservices/mall/service/product/rpc/product"

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

func (l *UpdateLogic) Update(in *product.UpdateRequest) (*product.UpdateResponse, error) {
	// product update
	_, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "Product not found!")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.ProductModel.Update(l.ctx, &model.Product{
		Id:     in.Id,
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Amount: in.Amount,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.UpdateResponse{}, nil
}
