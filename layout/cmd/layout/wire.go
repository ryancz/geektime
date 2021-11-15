// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"geektime/layout/internal/biz"
	"geektime/layout/internal/conf"
	"geektime/layout/internal/data"
	"geektime/layout/internal/data/db"
	"geektime/layout/internal/server"
	"geektime/layout/internal/service"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func initApp(cs *conf.Server, cd *conf.Data) *grpc.Server {
	panic(wire.Build(biz.ProviderSet, data.ProviderSet, db.ProviderSet, server.ProviderSet, service.ProviderSet))
}
