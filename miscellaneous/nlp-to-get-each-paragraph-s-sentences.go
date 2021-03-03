package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s url\n", os.Args[0])
		os.Exit(0)
	}

	url := os.Args[1]

	// perform a simple sanity check
	if url == "" {
		fmt.Println("URL address cannot be empty!")
		os.Exit(0)
	}

	fmt.Println("Grabbing text from : ", url)

	allTextData := GetArticleText(url)

	fmt.Println("All : ", allTextData)
	fmt.Println("-----------------------------------------------------")

	paragraphs := strings.Split(allTextData, "\n")
	//fmt.Println("Paragraph 0 : ", paragraph[0])
	//fmt.Println("Paragraph 1 : ", paragraph[1])
	//fmt.Println("Paragraph 2 : ", paragraph[2])

	// Create a new document for each paragraph
	for k, v := range paragraphs {
		fmt.Println("Processing paragraph ", k, " : ")
		doc, err := prose.NewDocument(v)
		if err != nil {
			log.Fatal(err)
		}

		// Iterate over the doc's sentences:
		for _, sent := range doc.Sentences() {
			fmt.Println("[" + sent.Text + "]")
		}
	}
}

func GetArticleText(url string) string {
	doc, err := goquery.NewDocument(url)

	if err != nil {
		panic(err)
	}

	allText := ""

	doc.Find(".story_paragraph p").Each(func(i int, s *goquery.Selection) {
		<---
		- modify
		here
		for other
		websites
		// For each item found, get the paragraph
		paragraph := s.Text()
		//fmt.Println("Paragraph >>>> ", paragraph)
		allText = allText + "\n" + paragraph

	})
	return allText
}

/*
 usage:
 	$ ./go-file-name.go https://www.dailyfx.com/forex/fundamental/dailybriefing/sessionbriefing/euro_open/2019/06/25/EURUSD-Uptrend-May-be-Accelerated-by-Powell-Commentary-US-Data.html
*/
