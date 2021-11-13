// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"geektime/layout/app/user/service/internal/biz"
	"geektime/layout/app/user/service/internal/data"
	"geektime/layout/app/user/service/internal/server"
	"geektime/layout/app/user/service/internal/service"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func initApp() (*grpc.Server, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}

