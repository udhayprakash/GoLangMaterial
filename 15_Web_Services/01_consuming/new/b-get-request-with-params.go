package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	name := "John Doe"
	occupation := "gardener"

	params := "name=" + url.QueryEscape(name) + "&" +
		"occupation=" + url.QueryEscape(occupation)
	path := fmt.Sprintf("https://httpbin.org/get?%s", params)

	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
