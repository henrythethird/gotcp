package main

import (
	"log"
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
		log.Fatalln(err.Error())
	}

	errChan := make(chan error)

	go func() {
		errChan <- client.Listen()
	}()

	tick := time.Tick(100 * time.Millisecond)

	for {
		select {
		case <-tick:
			select {
			case e := <-errChan:
				log.Fatalln(e.Error())
			default:
				events := handler.Handle()

				if len(events) > 0 {
					log.Println(events)
				}
			}

			client.Send("hello :-)")
		}
	}
}
