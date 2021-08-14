package main

import (
    "fmt"
    "log"
    "net/http"
    "net/url"
    "os"
)

func main() {
    req, err := http.NewRequest("GET","http://api.themoviedb.org/3/tv/popular", nil)
    if err != nil {
        log.Print(err)
        os.Exit(1)
    }

    // if you appending to existing query this works fine 
    q := req.URL.Query()
    q.Add("api_key", "key_from_environment_or_flag")
    q.Add("another_thing", "foo & bar")

    // or you can create new url.Values struct and encode that like so
    q := url.Values{}
    q.Add("api_key", "key_from_environment_or_flag")
    q.Add("another_thing", "foo & bar")

    req.URL.RawQuery = q.Encode()

    fmt.Println(req.URL.String())
    // Output:
    // http://api.themoviedb.org/3/tv/popularanother_thing=foo+%26+bar&api_key=key_from_environment_or_flag
}