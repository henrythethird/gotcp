package gotcp

import (
	"log"
	"net"

	"github.com/henrythethird/gotcp/event"
)

// Server is the server side of client-server communication
type Server struct {
	handler *event.Handler
	clients []*Client
}

// NewServer instantiates a new TCP server
func NewServer(handler *event.Handler) *Server {
	return &Server{
		handler: handler,
		clients: make([]*Client, 32),
	}
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
			cliSocket,
		)

		s.clients = append(s.clients, client)

		go client.Listen()
	}
}
