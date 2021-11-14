package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// go get -u github.com/PuerkitoBio/goquery

func main() {
	// Request html page
	res, err := http.Get("http://journaldev.com")
	CheckError(err)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// To display on the terminal
	// n, err := io.Copy(os.Stdout, res.Body)
	// CheckError(err)
	// log.Println("Number of bytes copied to STDOUT:", n)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	CheckError(err)

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		txt := s.Text()
		fmt.Printf("Article %d: %s\n\n", i, txt)
	})
}

func CheckError(er error) {
	if er != nil {
		log.Fatal(er)
	}
}
