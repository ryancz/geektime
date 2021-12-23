package main

import (
	"bytes"
	"encoding/binary"
)

type Buffer struct {
	data []byte
}

func NewBuffer() *Buffer {
	return &Buffer{}
}

func (b *Buffer) Fill(d []byte) {
	b.data = append(b.data, d...)
}

func (b *Buffer) Decode() (*Message, error) {
	if len(b.data) < MinHeaderLen {
		return nil, nil
	}

	var msg Message

	// Len 4 bytes
	bb := bytes.NewBuffer(b.data[0:4])
	err := binary.Read(bb, binary.BigEndian, &msg.Len)
	if err != nil {
		return nil, err
	}

	// HeaderLen 2 bytes
	bb = bytes.NewBuffer(b.data[4:6])
	err = binary.Read(bb, binary.BigEndian, &msg.HeaderLen)
	if err != nil {
		return nil, err
	}

	// Version 2 bytes
	bb = bytes.NewBuffer(b.data[6:8])
	err = binary.Read(bb, binary.BigEndian, &msg.Version)
	if err != nil {
		return nil, err
	}

	// Operation 4 bytes
	bb = bytes.NewBuffer(b.data[8:12])
	err = binary.Read(bb, binary.BigEndian, &msg.Operation)
	if err != nil {
		return nil, err
	}

	// SequenceId 4 bytes
	bb = bytes.NewBuffer(b.data[12:16])
	err = binary.Read(bb, binary.BigEndian, &msg.SeqId)
	if err != nil {
		return nil, err
	}

	// Body Len-16 bytes
	msg.Body = append(msg.Body, b.data[16:msg.Len]...)

	// remove message data
	b.data = b.data[msg.Len:]

	return &msg, nil
}
