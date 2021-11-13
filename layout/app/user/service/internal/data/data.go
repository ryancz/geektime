package data

import (
	"context"
	"fmt"
	"log"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ProviderSet = wire.NewSet(NewData, NewDb, NewUserRepo)

func NewDb() *gorm.DB {
	user := "root"
	password := "123456"
	host := "localhost"
	port := 3306
	name := "test"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user, password, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	return db
}

type Data struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) *Data {
	return &Data {
		db: db,
	}
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

type User struct {
	Id int64
	Username string
	Password string
}

func (d *Data) GetUser(ctx context.Context, id int64) (*User, error) {
	var user User
	if err := d.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *Data) CreateUser(ctx context.Context, user *User) (*User, error) {
	if err := d.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}