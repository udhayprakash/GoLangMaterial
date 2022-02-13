package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler2(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	// Set up a /hello resource handler
	http.HandleFunc("/hello", helloHandler2)

	// Listen to port 8080 and wait
	// log.Fatal(http.ListenAndServe("localhost:8080", nil))

	// Listen to HTTPS connections on port 8443 and wait
	log.Fatal(http.ListenAndServeTLS("localhost:8443", "cert.pem", "key.pem", nil))
}

/*
Command to create certificate for localhost:

openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -out cert.pem -keyout key.pem -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost" -addext "subjectAltName = DNS:localhost"

Generally, openssl will be installed alongside tools like git, etc
Also, there is a web tool for generating it - https://certificatetools.com/
*/
