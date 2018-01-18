package main

import (
	"fmt"
	"time"

	"github.com/henrythethird/gotcp"
	"github.com/henrythethird/gotcp/event"
)

func main() {
	handler := event.NewHandler()

	client := gotcp.NewClient(
		event.NewDispatcher(handler),
	)

	err := make(chan error)

	go func() {
		err <- client.ConnectAndListen("127.0.0.1:8888")
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

		time.Sleep(1000000000)

		client.Send("hello :-)")
	}
}
