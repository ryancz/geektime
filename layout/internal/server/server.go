package server

import (
	v1 "geektime/layout/api/user/v1"
	"geektime/layout/internal/conf"
	"geektime/layout/internal/service"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

var ProviderSet = wire.NewSet(NewGrpcServer)

func NewGrpcServer(c *conf.Server, us *service.UserService) *grpc.Server {
	var opts []grpc.ServerOption
	if c.Grpc.Timeout > 0 {
		opts = append(opts, grpc.ConnectionTimeout(c.Grpc.Timeout))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServiceServer(srv, us)
	return srv
}
