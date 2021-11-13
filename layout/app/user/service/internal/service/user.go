package service

import (
	"context"
	v1 "geektime/layout/api/user/service/v1"
	"geektime/layout/app/user/service/internal/biz"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	uc  *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{
		uc: uc,
	}
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	user, err := s.uc.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{
		Id: user.Id,
		Username: user.Username,
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	user, err := s.uc.CreateUser(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{
		Id: user.Id,
		Username: user.Username,
	}, nil
}