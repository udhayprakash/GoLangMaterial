package main

import (
	"net/http"
)

func Employees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// NOTE : put jsonStr in a FOR loop to read forever from database and encode the data with
	// json.Marshal() function to produce JSON stream
	// see https://www.socketloop.com/tutorials/golang-convert-csv-data-to-json-format-and-save-to-file

	// for this example, we will use static JSON string

	jsonStr := `[{"Name":"Dennis","Age":16,"Job":"CEO"},
   {"Name":"Eva","Age":34,"Job":"CFO"},
   {"Name":"Paul","Age":28,"Job":"COO"}]`

	// output to web or client
	w.Write([]byte(jsonStr))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Employees)

	http.ListenAndServe(":8080", mux)
}
