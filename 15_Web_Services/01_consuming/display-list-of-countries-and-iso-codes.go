package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// NOTE : struct doesn't work here due to the nature
// of the JSON data from country.io
// it has no field identifiers

//type Country struct {
//	ISO string
//}

func main() {

	resp, err := http.Get("http://country.io/names.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	jsonCountriesData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(string(jsonData))

	countriesMap := make(map[string]interface{})

	//NOTE : unmarshal to struct will throw this error message
	//        json: cannot unmarshal object into Go value of type

	// Decode JSON into our map
	err = json.Unmarshal([]byte(jsonCountriesData), &countriesMap)

	if err != nil {
		println(err)
		return
	}

	for iso2, name := range countriesMap {
		fmt.Println("ISO2 code:", iso2, "Country name :", name)
	}

}
