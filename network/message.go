package main

const MinHeaderLen = 16

type Message struct {
	Len       uint32
	HeaderLen uint16
	Version   uint16
	Operation uint32
	SeqId     uint32
	Body      []byte
}
