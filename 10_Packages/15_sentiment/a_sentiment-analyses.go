package main

import (
	"log"

	"github.com/cdipaolo/sentiment"
)

// go get -u github.com/cdipaolo/sentiment

var analysis *sentiment.Analysis

func sentimentCheck(text string) {
	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}
	analysis = model.SentimentAnalysis(text, sentiment.English)
	if analysis.Score == 1 {
		log.Printf("%s - Score of %d = Positive Sentiment\n", text, analysis.Score)
	} else {
		log.Printf("%s - Score of %d = Negative Sentiment\n", text, analysis.Score)
	}
}

func main() {
	// Negative Example
	sentimentCheck("Your mother is an awful lady")
	// sentimentCheck("Sua mãe é uma senhora horrível")

	// Positive Example
	sentimentCheck("Your mother is a lovely lady")

}
