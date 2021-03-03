package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", ExampleHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"ip": GetIP(r),
	})
	w.Write(resp)
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
