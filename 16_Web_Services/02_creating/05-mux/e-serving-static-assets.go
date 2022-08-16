package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	http.HandleFunc("/files", streamFiles)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func streamFiles(w http.ResponseWriter, r *http.Request) {
	// To get current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	// To list files & folders in current directory
	_files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range _files {
		fmt.Println(f.Name())
	}

	// To read files in the "files" folder
	files, err := ioutil.ReadDir("files")
	if err != nil {
		log.Fatal(err)
	}
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	for _, f := range files {
		buf, err := ioutil.ReadFile(path.Join("./files", f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		w.Write(buf)
		flusher.Flush()
	}
}
