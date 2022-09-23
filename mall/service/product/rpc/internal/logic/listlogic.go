package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/product/rpc/internal/svc"
	"go-zero_microservices/mall/service/product/rpc/product"

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

func (l *ListLogic) List(in *product.ListRequest) (*product.ListResponse, error) {
	// product list
	var (
		list   []*product.DetailResponse
		single *product.DetailResponse
	)
	products, err := l.svcCtx.ProductModel.FindAll(l.ctx, in.Page, in.Size, in.Name)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	for _, v := range products {
		single = &product.DetailResponse{
			Id:     v.Id,
			Name:   v.Name,
			Desc:   v.Desc,
			Stock:  v.Stock,
			Amount: v.Amount,
			Status: v.Status,
		}
		list = append(list, single)
	}

	return &product.ListResponse{
		List: list,
	}, nil
}
