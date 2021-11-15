package db

import (
	"log"

	"geektime/layout/internal/conf"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ProviderSet = wire.NewSet(NewGormDb, NewClient)

func NewGormDb(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Db.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	return db
}

type Client struct {
	User *UserClient
}

func NewClient(db *gorm.DB) *Client {
	return &Client{
		User: NewUserClient(db),
	}
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
