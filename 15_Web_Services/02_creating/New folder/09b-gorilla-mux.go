package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	route()
}
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World !!!"))
}
func ProcessCrons(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Crons route"))
	/* How to write multiple routes here */
	//r.PathPrefix("/crons/add_data.php").HandlerFunc(AddDataCron).Methods("GET", "POST")
	//r.PathPrefix("/crons/remove_data.php").HandlerFunc(RemoveDataCron).Methods("GET", "POST")
}
func route() {
	r := mux.NewRouter()
	r.PathPrefix("/hello").HandlerFunc(hello).Methods("GET", "POST")
	r.PathPrefix("/crons/").HandlerFunc(ProcessCrons).Methods("GET", "POST")
	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  5 * time.Second,
		Handler:      handlers.CompressHandler(r),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}