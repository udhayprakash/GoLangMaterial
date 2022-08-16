package main

import (
	"log"
	"net/http"
	"time"
)

type foo struct{}

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("New request")
	for i := 0; i < 10; i++ {
		select {
		case <-r.Context().Done():
			log.Print("Aborted")
			return
		case <-time.After(1 * time.Second):
			log.Print("Tick")
		}
		w.Write([]byte(".\n"))
	}
	w.Write([]byte("hello world\n"))
	log.Print("Completed")
}

func main() {
	fooHandler := foo{}
	timeoutHandler := http.TimeoutHandler(fooHandler, 5*time.Second, "Too slow!\n")
	http.Handle("/foo", timeoutHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

To test:
curl localhost:8080/foo