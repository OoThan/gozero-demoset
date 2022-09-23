package logic

import (
	"context"
	"go-zero_microservices/mall/service/product/rpc/product"

	"go-zero_microservices/mall/service/product/api/internal/svc"
	"go-zero_microservices/mall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListRequest) (*types.ListResponse, error) {
	// product list
	var productList = make([]types.DetailResponse, 0)
	res, err := l.svcCtx.ProductRpc.List(l.ctx, &product.ListRequest{
		Page: req.Page,
		Size: req.Size,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	for _, v := range res.List {
		productList = append(productList, types.DetailResponse{
			Id:     v.Id,
			Name:   v.Name,
			Desc:   v.Desc,
			Stock:  v.Stock,
			Amount: v.Amount,
			Status: v.Status,
		})
	}

	return &types.ListResponse{
		List: productList,
	}, nil
}
