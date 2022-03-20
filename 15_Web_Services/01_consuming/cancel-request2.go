package main

import (
	"context"
	"net/http"
	"time"
)

func main() {
	cx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequest("GET", "http://google.com", nil)
	req = req.WithContext(cx)
	ch := make(chan error)

	go func() {
		_, err := http.DefaultClient.Do(req)
		select {
		case <-cx.Done():
			// Already timedout
		default:
			ch <- err
		}
	}()

	// Simulating user cancel request
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()
	select {
	case err := <-ch:
		if err != nil {
			// HTTP error
			panic(err)
		}
		print("no error")
	case <-cx.Done():
		panic(cx.Err())
	}

}
