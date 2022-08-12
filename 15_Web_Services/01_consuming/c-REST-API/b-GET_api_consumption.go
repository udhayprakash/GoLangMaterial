package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Page       int
	Per_Page   int
	Total      int
	TotalPages int
	Data       []Article
}

type Article struct {
	Title        string
	Url          string
	Author       string
	Num_Comments int
	Story_Id     int
	Story_Title  string
	Story_Url    string
	Parent_Id    int
	Created_At   time.Time
}

func main() {
	author := "apaga"
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/articles?author=%s", author)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	jsonString := buf.String()
	fmt.Println("jsonString:", jsonString)
	var result Result
	json.Unmarshal([]byte(jsonString), &result)
	fmt.Println("result.Page:", result.Page)
	fmt.Println(result.Data)
	for _, value := range result.Data {
		fmt.Println("value.Title:", value.Title)
		fmt.Println("value.Url:", value.Url)
		fmt.Println("value.Author:", value.Author)
		fmt.Println("value.Num_Comments:", value.Num_Comments)
	}
	// fmt.Println("page:", jsonString["page"])
	// fmt.Println("per_page:", jsonString["per_page"])
	// json.Unmarshal(resp.Body)
	// return []string{result}
}
