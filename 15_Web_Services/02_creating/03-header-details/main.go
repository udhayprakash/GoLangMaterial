package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Routing
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", requestHeaders)
	http.HandleFunc("/request-details", requestDetails)

	log.Println("Starting server at 8090 ...")

	// Starting server - default is localhost
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func requestHeaders(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%25v: %v\n", name, h)
		}
	}
}

func requestDetails(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "RequestURI: %v\n", req.RequestURI)
	fmt.Fprintf(resp, "req.URL   : %v\n", req.URL)
	fmt.Fprintf(resp, "String()  : %v\n", req.URL.String())
	fmt.Fprintf(resp, "Path      : %v\n", req.URL.Path)
	fmt.Fprintf(resp, "RequestURI: %v\n", req.URL.RequestURI())
	fmt.Fprintf(resp, "Opaque    : %v\n", req.URL.Opaque)
	fmt.Fprintf(resp, "RawPath   : %v\n", req.URL.RawPath)
	fmt.Fprintf(resp, "RawQuery  : %v\n", req.URL.RawQuery)
	fmt.Fprintf(resp, "User      : %v\n", req.URL.User)
	fmt.Fprintf(resp, "IsAbs()   : %v\n", req.URL.IsAbs())
	fmt.Fprintf(resp, "Scheme    : %v\n", req.URL.Scheme)
}
