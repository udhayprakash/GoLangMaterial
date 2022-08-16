package main

import (
	"fmt"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func main() {
	// creating a single router
	router := bunrouter.New()

	// Adding routes
	router.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		// req embeds *http.Request and has all the same fields and methods
		fmt.Println(req.Method, req.Route(), req.Params().Map())
		return nil
	})
	fmt.Println(router)
}
