package main

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

// go get -u github.com/jarcoal/httpmock

func TestScrapeTitle2(t *testing.T) {
	fmt.Println("TestScrapeTitle")
	myMockPage := `<!DOCTYPE html><html lang=en-US class=no-js><head><title>What does the Otter say?</title></head></html>`

	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://jonathanmh.com",
		httpmock.NewStringResponder(200, myMockPage))

	siteTitle := ScrapeTitle("https://jonathanmh.com")
	httpmock.DeactivateAndReset()

	if siteTitle != "What does the Otter say?" {
		t.Error("<title> was not correctly extracted")
	}

	fmt.Println(siteTitle)
}

/*
myMockPage :=... 				sets up our example response, a piece of plain text that our function will
								parse into a HTML and look for the title
httpmock.Activate() 			activates the mocking, before this no requests can be intercepted
httpmock.RegisterResponder() 	defines the METHOD and the URL, so GET or POST and an address at which we fake
								an http response
httpmock.NewStringResponder 	will need a status code and a string to respond with instead of what actually
								lives at that URL
httpmock.DeactivateAndReset() 	stops mocking responses for the rest of the test


*/
