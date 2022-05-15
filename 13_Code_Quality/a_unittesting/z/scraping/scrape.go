package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeTitle(websiteURL string) (title string) {
	doc, err := goquery.NewDocument(websiteURL)
	if err != nil {
		log.Fatalln("could not load website for scraping")
	}
	siteTitle := doc.Find("title").Contents().Text()
	return siteTitle
}
