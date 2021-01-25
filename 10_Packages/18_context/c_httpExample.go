package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))

}

func main() {

	srv := http.Server{Addr: ":8081"}

	idleConnsClosed := make(chan struct{})

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		log.Print("HTTP server Shutdown successful")
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed

}
