package gotcp

import (
	"log"
	"net"

	"github.com/henrythethird/gotcp/event"
)

// Server is the server side of client-server communication
type Server struct {
	handler *event.Handler
}

// NewServer instantiates a new TCP server
func NewServer(handler *event.Handler) *Server {
	return &Server{handler}
}

// Listen starts listening for client communication
func (s *Server) Listen(address string) error {
	listener, err := net.Listen("tcp", address)

	if err != nil {
		return err
	}

	defer listener.Close()

	for {
		cliSocket, err := listener.Accept()

		if err != nil {
			log.Println("Client failed to connect:", err.Error())
			continue
		}

		defer cliSocket.Close()

		client := NewClient(
			event.NewDispatcher(s.handler),
		)

		go client.Listen(cliSocket)
	}
}
