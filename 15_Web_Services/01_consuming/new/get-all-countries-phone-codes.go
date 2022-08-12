package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://country.io/phone.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	jsonPhoneCodesData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonPhoneCodesData))

	phoneMap := make(map[string]interface{})
	// phoneMap := make(map[string]string)

	// Decode JSON into our map
	err = json.Unmarshal([]byte(jsonPhoneCodesData), &phoneMap)

	if err != nil {
		println(err)
		return
	}

	for iso2, phoneCode := range phoneMap {
		fmt.Println("ISO2 code:", iso2, "Phone code ", phoneCode)
	}
}
