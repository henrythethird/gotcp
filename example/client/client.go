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
		nil,
	)

	err := client.Connect("127.0.0.1:8888")

	if err != nil {
		fmt.Println(err.Error())
	}

	errChan := make(chan error)

	go func() {
		errChan <- client.Listen()
	}()

	for {
		select {
		case e := <-errChan:
			fmt.Println(e.Error())
			return
		default:
			events := handler.Handle()

			if len(events) > 0 {
				fmt.Println(events)
			}
		}

		client.Send("hello :-)")
		time.Sleep(100000000)
	}
}
