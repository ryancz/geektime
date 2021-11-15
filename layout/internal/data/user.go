package data

import (
	"context"
	"geektime/layout/internal/biz"
	"geektime/layout/internal/data/db"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{data: data}
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	po, err := r.data.db.User.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return poToDoUser(po), nil
}

func poToDoUser(po *db.User) *biz.User {
	return &biz.User{
		Id:       po.Id,
		Username: po.Username,
		Password: po.Password,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	po, err := r.data.db.User.CreateUser(ctx, &db.User{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		return nil, err
	}
	return poToDoUser(po), nil
}
