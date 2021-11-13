package server

import (
	v1 "geektime/layout/api/user/service/v1"
	"geektime/layout/app/user/service/internal/service"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

var ProviderSet = wire.NewSet(NewGrpcServer)

func NewGrpcServer(us *service.UserService) *grpc.Server {
	srv := grpc.NewServer()
	v1.RegisterUserServer(srv, us)
	return srv
}
