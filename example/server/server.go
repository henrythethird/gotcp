package main

import (
	"fmt"

	"github.com/henrythethird/gotcp"
	"github.com/henrythethird/gotcp/event"
)

func main() {
	handler := event.NewHandler()
	server := gotcp.NewServer(handler)

	err := make(chan error)

	go func() {
		err <- server.Listen("127.0.0.1:8888")
	}()

	for {
		select {
		case e := <-err:
			fmt.Println(e.Error())
			return
		default:
			events := handler.Handle()

			if len(events) > 0 {
				fmt.Println(events)
			}
		}
	}
}
