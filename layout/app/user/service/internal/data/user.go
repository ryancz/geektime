package data

import (
	"context"
	"geektime/layout/app/user/service/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	po, err := r.data.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id: po.Id,
		Username: po.Username,
	}, nil
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	po, err := r.data.CreateUser(ctx, &User{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id: po.Id,
		Username: po.Username,
	}, nil
}