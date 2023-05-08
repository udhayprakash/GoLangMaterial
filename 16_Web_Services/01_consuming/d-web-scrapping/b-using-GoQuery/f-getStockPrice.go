package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	symbol = "AAPL"
)

func main() {
	res, err := http.Get("https://finance.yahoo.com/quote/" + symbol)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code %d: %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	marketPrice := doc.Find("[data-field=regularMarketPrice][data-symbol=" + symbol + "]").Text()

	if marketPrice == "" {
		log.Fatalf("cannot access market price")
	}

	fmt.Printf("%s price: %s\n", symbol, marketPrice)
}