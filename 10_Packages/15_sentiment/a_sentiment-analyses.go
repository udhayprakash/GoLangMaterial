package main

import (
	"github.com/cdipaolo/sentiment"
	"log"
)

func main() {

	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}

	var analysis *sentiment.Analysis
	var text string

	// Negative Example
	text = "Your mother is an awful lady"
	analysis = model.SentimentAnalysis(text, sentiment.English)
	if analysis.Score == 1 {
		log.Printf("%s - Score of %d = Positive Sentiment\n", text, analysis.Score)
	} else {
		log.Printf("%s - Score of %d = Negative Sentiment\n", text, analysis.Score)
	}

	// Positive Example
	text = "Your mother is a lovely lady"
	analysis = model.SentimentAnalysis(text, sentiment.English)
	if analysis.Score == 1 {
		log.Printf("%s - Score of %d = Positive Sentiment\n", text, analysis.Score)
	} else {
		log.Printf("%s - Score of %d = Negative Sentiment\n", text, analysis.Score)
	}
}
