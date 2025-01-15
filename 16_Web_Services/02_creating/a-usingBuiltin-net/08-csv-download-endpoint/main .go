package main

import (
	"encoding/csv"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveDummyCSV)
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}

func serveDummyCSV(w http.ResponseWriter, req *http.Request) {
	items := [][]string{
		{"UserID", "FullName", "Email"},           // Header
		{"1", "Jack Johnson", "jack@hotmail.com"}, // Items
		{"2", "Jill Smith", "jill@hotmail.com"},
		{"3", "James Murphy", "james@hotmail.com"},
	}

	// Set our headers so browser will download the file
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=users.csv")
	// Create a CSV writer using our HTTP response writer as our io.Writer
	wr := csv.NewWriter(w)
	// Write all items and deal with errors
	if err := wr.WriteAll(items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
