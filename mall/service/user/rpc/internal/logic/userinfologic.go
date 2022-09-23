package logic

import (
	"context"
	"go-zero_microservices/mall/service/user/model"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/user/rpc/internal/svc"
	"go-zero_microservices/mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// check user
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "user doesn't exist.")
		}

		return nil, status.Error(500, err.Error())
	}

	return &user.UserInfoResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: int32(res.Gender),
		Mobile: res.Mobile,
	}, nil
}
