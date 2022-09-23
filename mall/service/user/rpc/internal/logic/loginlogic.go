package logic

import (
	"context"
	"go-zero_microservices/mall/service/user/model"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/user/rpc/internal/svc"
	"go-zero_microservices/mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "User does not exist")
		}
		return nil, status.Error(500, err.Error())
	}

	// check password

	return &user.LoginResponse{}, nil
}
