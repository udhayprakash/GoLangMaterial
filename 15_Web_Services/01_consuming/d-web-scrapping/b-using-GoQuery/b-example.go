package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	CheckError(err)

	dhead := doc.Find("head")
	dTitle := dhead.Find("title")
	fmt.Printf("title text:%s\n", dTitle.Text())

	html, _ := dTitle.Html()
	fmt.Printf("title html:%s\n", html)

	metaArr := dhead.Find("meta")
	for i := 0; i < metaArr.Length(); i++ {
		d, _ := metaArr.Eq(i).Attr("name")
		fmt.Println(d)
	}
	doc.Find("div.wrapper .container .col-lg-9").Each(func(i int, cs *goquery.Selection) {
		d, _ := cs.Attr("class")
		fmt.Println(d)
	})

}

// Output:
// col-lg-9 col-md-8 col-sm-7

func CheckError(er error) {
	if er != nil {
		log.Fatal(er)
	}
}
