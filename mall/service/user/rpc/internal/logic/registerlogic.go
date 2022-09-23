package logic

import (
	"context"
	"go-zero_microservices/mall/common/cryptx"
	"go-zero_microservices/mall/service/user/model"
	"google.golang.org/grpc/status"

	"go-zero_microservices/mall/service/user/rpc/internal/svc"
	"go-zero_microservices/mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// determine whether the phone number is registered
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Password)
	if err != nil {
		return nil, status.Error(100, "user already exists")
	}

	if err == model.ErrNotFound {
		newUser := model.User{
			Name:     in.Name,
			Gender:   int64(in.Gender),
			Mobile:   in.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUser.Id, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.RegisterResponse{
			Id:     newUser.Id,
			Name:   newUser.Name,
			Gender: int32(newUser.Gender),
			Mobile: newUser.Mobile,
		}, nil
	}

	return nil, status.Error(500, err.Error())
}
