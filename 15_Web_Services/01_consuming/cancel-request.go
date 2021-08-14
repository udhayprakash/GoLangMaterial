package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	req, _ := http.NewRequest("GET", "http://google.com", nil)
	tr := &http.Transport{} // TODO: copy defaults from http.DefaultTransport
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)
	go func() {
		resp, err := client.Do(req)
		// handle response ...
		_ = resp
		c <- err
	}()

	// Simulating user cancel request channel
	user := make(chan struct{}, 0)
	go func() {
		time.Sleep(100 * time.Millisecond)
		user <- struct{}{}
	}()

	for {
		select {
		case <-user:
			log.Println("Cancelling request")
			tr.CancelRequest(req)
		case err := <-c:
			log.Println("Client finished:", err)
			return
		}
	}
}
