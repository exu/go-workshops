package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pair"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

var (
	node string
)

func newSocket() mangos.Socket {
	socket, err := pair.NewSocket()
	if err != nil {
		log.Fatalf("Node %s: Cannot create socket: %s\n", node, err.Error())
	}

	socket.AddTransport(ipc.NewTransport())
	socket.AddTransport(tcp.NewTransport())

	socket.SetOption(mangos.OptionRecvDeadline, 10*time.Second)
	return socket
}

func send(socket mangos.Socket, message string) {
	log.Printf("Node %s sends %s\n", node, message)
	err := socket.Send([]byte(message))
	if err != nil {
		log.Fatalf("Node %s failed to send '%s': %s\n", node, message, err.Error())
	}
}

func receive(socket mangos.Socket) string {

	bytes, err := socket.Recv()
	if err != nil {
		log.Fatalf("Node %s failed receiving a message: %s\n", node, err.Error())
	}
	message := string(bytes)
	log.Printf("Node %s received %s\n", node, message)
	return message
}

func runNode(url string) {
	socket := newSocket()
	err := socket.Listen(url)
	if err != nil {
		log.Printf("Node %s cannot listen on socket '%s': %s\nTrying to dial instead\n", node, url, err.Error())
		err = socket.Dial(url)
		if err != nil {
			log.Fatalf("Node %s can neither listen nor dial on socket '%s': %s\n", node, url, err.Error())
		}
	}

	defer socket.Close()
	for i := 0; i < 3; i++ {
		send(socket, fmt.Sprintf("message %d from node %s.", i, node))
		_ = receive(socket)
		time.Sleep(1 * time.Second)
	}
	log.Printf("Node %s: Done.\n", node)
}

func main() {
	if len(os.Args) <= 2 {
		log.Printf("Usage: %s 0|1 <url>\n", os.Args[0])
	} else {
		node = os.Args[1]
		runNode(os.Args[2])
	}
}
