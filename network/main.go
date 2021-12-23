package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("listen failed: %v", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept failed %v\n", err)
			break
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	b := NewBuffer()
	for {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if n > 0 {
			b.Fill(data)
			msg, err := b.Decode()
			if err != nil {
				log.Printf("decode message failed: %v\n", err)
				break
			}

			if msg != nil {
				log.Printf("message: %v\n", msg)
			}
		}

		if err != nil {
			log.Printf("read failed: %v", err)
			break
		}
	}
}
