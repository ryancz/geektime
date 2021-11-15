package db

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserClient struct {
	db *gorm.DB
}

func NewUserClient(db *gorm.DB) *UserClient {
	return &UserClient{db: db}
}

func (c *UserClient) GetUser(ctx context.Context, id int64) (*User, error) {
	var user User
	if err := c.db.First(&user, id).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("get user %d error", id))
	}
	return &user, nil
}

func (c *UserClient) CreateUser(ctx context.Context, user *User) (*User, error) {
	if err := c.db.Create(user).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("create user %s error", user.Username))
	}
	return user, nil
}
