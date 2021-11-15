package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"

	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"geektime/layout/internal/conf"
)

func main() {
	bc, err := loadConfig("../../config/config.yaml")
	if err != nil {
		log.Fatalf("load config error: %v", err)
	}

	gs := initApp(&bc.Server, &bc.Data)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return runSignal(ctx)
	})
	g.Go(func() error {
		return runGrpcServer(ctx, gs, bc.Server.Grpc.Addr)
	})
	if err = g.Wait(); err != nil {
		log.Println("errgroup wait error:", err)
	}
}

func loadConfig(filename string) (*conf.Bootstrap, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("os read file %s error", filename))
	}

	var bc conf.Bootstrap
	if err = yaml.Unmarshal(data, &bc); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unmarshal content of file %s to conf.Bootstrap error", filename))
	}
	return &bc, nil
}

func runGrpcServer(ctx context.Context, gs *grpc.Server, addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "listen tcp %s error")
	}

	go func() {
		<-ctx.Done()
		gs.GracefulStop()
	}()

	return gs.Serve(l)
}

func runSignal(ctx context.Context) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sig:
		return errors.New("received interrupt signal")
	case <-ctx.Done():
		return ctx.Err()
	}
}
