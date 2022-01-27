package main

import (
	"net/http"

	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func main() {
	// creating a single router
	// router := bunrouter.New()
	router := bunrouter.New(
		bunrouter.WithMiddleware(reqlog.NewMiddleware()), // to log requests
	)
	/*
		bunrouter.New accepts a couple of options to customize the router:
			- WithNotFoundHandler(handler bunrouter.HandlerFunc)
				- overrides the handler that is called when there are no matching routes.

			- WithMethodNotAllowedHandler(handler bunrouter.HandlerFunc)
				- overrides the handler that is called when a route matches, but the route
				  does not have a handler for the requested method.

			- WithMiddleware(middleware bunrouter.MiddlewareFunc)
				- adds the middleware to the router's stack of middlewares.
				- Router will apply the middleware to all route handlers including
				  NotFoundHandler and MethodNotAllowedHandler.

	*/
	// Attach routes to specific HTTP methods
	router.POST("/users", createUserHandler)
	router.GET("/users/:id", showUserHandler)
	router.PUT("/users/:id", updateUserHandler)
	router.DELETE("/users/:id", deleteUserHandler)
}

// Handlers
/*
bunrouter.HandlerFunc is a slightly enhanced version of http.HandlerFunc
that accepts bunrouter.Request and returns an error that you can handle with middlewares:
*/
func createUserHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}

func showUserHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}
func updateUserHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}

func deleteUserHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}


/*
Compat Handler
	router := bunrouter.New().Compat()

	router.GET("/", func(w http.ResponseWriter, req *http.Request) {
		params := bunrouter.ParamsFromContext(req.Context())
		fmt.Println(params.Route(), params.Map())
	})

Verbose Handler
	router := bunrouter.New().Verbose()

	router.GET("/", func(w http.ResponseWriter, req *http.Request, ps bunrouter.Params) {
		fmt.Println(params.Route(), params.Map(), params.ByName("param"))
	})

*/

// Ref - https://bunrouter.uptrace.dev/guide/handlers.html#panic-recovering