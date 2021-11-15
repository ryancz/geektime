package conf

import "time"

type Bootstrap struct {
	Server Server
	Data   Data
}

type Server struct {
	Grpc Grpc
}

type Grpc struct {
	Addr    string
	Timeout time.Duration
}

type Data struct {
	Db Db
}

type Db struct {
	Source string
}
