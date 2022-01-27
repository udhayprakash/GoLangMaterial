package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT,CREATE,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "DELETE":
		Delete(w, r)
	case "POST":
		Create(w, r)
	case "PUT":
		Update(w, r)
	default: //GET
		Select(w, r)
	}
}

func Select(w http.ResponseWriter, r *http.Request) {
	route, val := getpath(r)
	data := map[string]string{"method": "GET", "route": route, "value": val}
	json.NewEncoder(w).Encode(data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	route, val := getpath(r)
	data := map[string]string{"method": "PUT", "route": route, "value": val}
	json.NewEncoder(w).Encode(data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	route, val := getpath(r)
	data := map[string]string{"method": "POST", "route": route, "value": val}
	json.NewEncoder(w).Encode(data)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	route, val := getpath(r)
	data := map[string]string{"method": "DELETE", "route": route, "value": val}
	json.NewEncoder(w).Encode(data)
}

func getpath(r *http.Request) (string, string) {
	path := strings.Split(r.URL.String(), "/")
	switch len(path) {
	case 3:
		return path[1], path[2]
	case 2:
		return path[1], ""
	default:
		return "", ""
	}
}