package main

import (
	"github.com/lzzjingcairensheng/delay-queue/sever"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":6789")
	if err != nil {
		panic(err.Error())
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		sever.Pool.Process(conn)
	}
}
