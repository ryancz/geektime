package main

import (
	"log"
	"net"
)

func main() {
	gs, err := initApp()
	if err != nil {
		log.Fatalf("init failed %v\n", err)
	}

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("listen tcp 1234 failed %v\n", err)
	}

	gs.Serve(l)
}
