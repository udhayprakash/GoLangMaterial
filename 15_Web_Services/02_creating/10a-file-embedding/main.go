package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
)

// content holds the static content for the web server.
//go:embed static/* 
var content embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(content)))
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	// Read files from the embedded file-system.
	entries, err := content.ReadDir("static")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		resp, err := http.Get("http://localhost:8080/static/" + e.Name())
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q: %s", e.Name(), body)
	}
}

/*
-- static/a.txt --
hello a

-- static/b.txt --
hello b
*/
