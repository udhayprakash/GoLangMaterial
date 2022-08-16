package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

type Tour struct {
	Name, Price string
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20) // 20 is some number

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
	return tours
}
func main() {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("Response type: %T\n", resp)

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	// fmt.Println(content)

	tours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}
}
