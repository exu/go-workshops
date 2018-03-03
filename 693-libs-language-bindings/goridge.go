package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/spiral/goridge"
)

type App struct{}

func (s *App) Hi(name string, r *string) error {
	*r = fmt.Sprintf("Hello, %s!", name)
	return nil
}

func (s *App) Ho(i map[string]string, r *string) error {
	*r = fmt.Sprintf("Hello, your 'array' has %d elements!", len(i))
	return nil
}

func main() {
	ln, err := net.Listen("tcp", ":6001")
	if err != nil {
		panic(err)
	}

	rpc.Register(new(App))

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeCodec(goridge.NewCodec(conn))
	}
}
