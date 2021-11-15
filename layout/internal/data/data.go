package data

import (
	"geektime/layout/internal/data/db"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type Data struct {
	db *db.Client
}

func NewData(db *db.Client) *Data {
	return &Data{db: db}
}
