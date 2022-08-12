package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// go get -u github.com/PuerkitoBio/goquery
func main() {
	html := `<html>
            <body>
                <h1 id="title">spring morning</h1>
                <p class="content1">
                Sleep in spring,
                I hear birds singing everywhere.
                The sound of wind and rain at night,
                How many flowers fall.
                </p>
            </body>
            </html>
            `
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	CheckError(err)

	dom.Find("p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

//The results are as follows
//Sleep in spring,
//I hear birds singing everywhere.
//The sound of wind and rain at night,
//How many flowers fall.

func CheckError(er error) {
	if er != nil {
		log.Fatal(er)
	}
}
