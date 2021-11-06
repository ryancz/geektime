package service

import (
	"geek/errors/dao"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func CreateUser(account, email string) (*dao.User, error) {
	if account == "" {
		return nil, errors.New("CreateUser failed: account empty")
	}
	if email == "" {
		return nil, errors.New("CreateUser failed: email empty")
	}

	user, err := dao.GetUserByAccount(account)
	if err == nil {
		return nil, errors.New("CreateUser failed: account exist")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.WithMessage(err, "CreateUser failed")
	}

	user = &dao.User{
		Account: account,
		Email: email,
	}
	return user, errors.WithMessage(user.Insert(), "CreateUser failed: ")
}
