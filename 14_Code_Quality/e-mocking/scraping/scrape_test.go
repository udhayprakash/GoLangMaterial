package main

import (
	"fmt"
	"testing"
)

func TestScrapeTitle(t *testing.T) {
	fmt.Println("TestScrapeTitle")
	siteTitle := ScrapeTitle("http://jonathanmh.com")
	if siteTitle != "Blog of JonathanMH" {
		t.Error("title incorrect")
	}
	fmt.Println(siteTitle)
}
