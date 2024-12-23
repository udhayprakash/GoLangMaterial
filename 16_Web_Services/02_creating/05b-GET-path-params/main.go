package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	http.HandleFunc("/hello/", hello)
	http.ListenAndServe("localhost:8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	var helloExp = regexp.MustCompile(`/hello/(?P<name>[a-zA-Z]+)/(?P<age>\d+)`)

	match := helloExp.FindStringSubmatch(req.URL.Path)
	if len(match) > 0 {
		result := make(map[string]string)
		for i, name := range helloExp.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		if _, err := strconv.Atoi(result["age"]); err == nil {
			fmt.Fprintf(w, "Hello, %v year old named %s!", result["age"], result["name"])
		} else {
			fmt.Fprintf(w, "Sorry, not accepted age!")
		}
	} else {
		fmt.Fprintf(w, "Wrong url\n")
	}
}

// http://localhost:8090/hello/John/25
// Hello, 25 year old named John!
