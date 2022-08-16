package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	blogTitles, err := GetLatestBlogTitles("https://golangcode.com")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Blog Titles:")
	fmt.Printf(blogTitles)
}

// GetLatestBlogTitles gets the latest blog title headings from the url
// given and returns them as a list.
func GetLatestBlogTitles(url string) (string, error) {

	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// Save each .post-title as a list
	titles := ""
	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
		titles += "- " + s.Text() + "\n"
	})
	return titles, nil
}
