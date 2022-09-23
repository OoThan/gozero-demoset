package logic

import (
	"context"
	"go-zero_microservices/mall/service/product/model"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/product/rpc/internal/svc"
	"go-zero_microservices/mall/service/product/rpc/product"

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

func (l *CreateLogic) Create(in *product.CreateRequest) (*product.CreateResponse, error) {
	// product create
	newProduct := model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Amount: in.Amount,
		Status: in.Status,
	}

	res, err := l.svcCtx.ProductModel.Insert(l.ctx, &newProduct)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.CreateResponse{
		Id: uint64(id),
	}, nil
}
