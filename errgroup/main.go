package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return runSignal(ctx)
	})
	g.Go(func() error {
		return runServer(ctx)
	})
	if err := g.Wait(); err != nil {
		fmt.Printf("%v\n", err)
	}
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

func runServer(ctx context.Context) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello errgroup"))
	})
	s := http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background()) // ctx已经Done了，这里不能传ctx吧
	}()

	return errors.Wrap(s.ListenAndServe(), "listen and serve error")
}
