package gotcp

import (
	"errors"
	"fmt"
	"net"

	"github.com/henrythethird/gotcp/event"
)

const bufferLength = 10000

// Client side of the TCP communication
type Client struct {
	parser     *parser
	packer     *packer
	dispatcher *event.Dispatcher
	socket     net.Conn
}

// NewClient instantiates a new client
func NewClient(dispatcher *event.Dispatcher) *Client {
	return &Client{
		parser:     newParser(),
		packer:     newPacker(),
		dispatcher: dispatcher,
	}
}

// ConnectAndListen connects to a server given its address and starts listeing
func (c *Client) ConnectAndListen(address string) error {
	socket, err := net.Dial("tcp", address)

	if err != nil {
		return err
	}

	defer socket.Close()

	return c.Listen(socket)
}

// Listen listens to messages from the server
func (c *Client) Listen(socket net.Conn) error {
	c.socket = socket

	buffer := make([]byte, bufferLength)

	c.emitEvent("connect", c)

	for {
		len, err := socket.Read(buffer)

		if err != nil {
			c.emitEvent("disconnect", c)
			return err
		}

		fmt.Println(
			c.parser.Parse(buffer, uint64(len)),
		)
	}
}

// Send transmits the payload to the server
func (c *Client) Send(payload string) error {
	if c.socket == nil {
		return errors.New("connection not initialized")
	}

	c.socket.Write(
		c.packer.Pack(payload),
	)

	return nil
}

func (c *Client) emitEvent(kind string, body interface{}) {
	c.dispatcher.Emit(
		event.Event{Kind: kind, Body: body},
	)
}
