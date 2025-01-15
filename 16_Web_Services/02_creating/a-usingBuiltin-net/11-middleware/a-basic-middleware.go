package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

/*
Middleware is an entity that intercepts the server's request/response life cycle.
  - Process the request before running business logic (authentication)
  - Modify the request to the next handler function (attaching payload)
  - Modify the response for the client
  - Logging.... and much more

*/
func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running before handler")
		w.Write([]byte("Hijacking Request "))
		originalHandler.Serve(w, r)
		fmt.Println("Running after handler")
	})
}

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("405 - Header Content-Type incorrect"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func setTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cookie here is a struct that represents an HTTP
		// cookie as sent in the Set-Cookie header of HTTP request
		cookie := http.Cookie{
			Name:  "Server Time (UTC)", // can be anything
			Value: strconv.Itoa(int(time.Now().Unix())),
			// 👆 converted time to string
		}
		// now set the cookie to response
		http.SetCookie(w, &cookie)
		handler.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "<h1>Hello!</h1>")
	fmt.Println("Executing the handler")
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	midHandler := http.HandlerFunc(helloHandler)
	http.HandleFunc("/helloWithMiddleware", middleware(midHandler))

	fmt.Println("Listening at localhost:3000")
	http.ListenAndServe(":3000", nil)
}

// // Chaining Middlewares
// type Middleware func(http.Handler) http.Handler
// type Chain []Middleware

// // returns a Slice of middlewares
// func CreateChain(middlewares ...Middleware) Chain {
//     var slice Chain
//     return append(slice, middlewares...)
// }

// func (c Chain) Then(originalHandler http.Handler) http.Handler {
//     if originalHandler == nil {
//         originalHandler = http.DefaultServeMux
//     }

//     for i := range c {
//         // Same as to m1(m2(m3(originalHandler)))
//         originalHandler = c[len(c) -1 -i](originalHandler)
//     }
//     return originalHandler
// }

// func main() {
//     myHandler := http.HandlerFunc(postHandler)
//     chain := CreateChain(filterContentType, setTimeCookie).Then(myHandler) // 👈
//     http.Handle("/", chain)
//     http.ListenAndServe(":8000", nil)
// }
