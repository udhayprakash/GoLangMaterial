package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	CheckError(err)
	// fmt.Println(doc.Html())            //. Html() gets the HTML content
	pTitle := doc.Find("title").Text() //Extract the content of title directly
	fmt.Printf("title:%v\n", pTitle)

	class := doc.Find("h2").Text()
	fmt.Printf("class:%v\n", class)

	doc.Find(".topics .topic").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".title a").Text()
		t := contentSelection.Find(".title a")
		log.Printf("the length;%d", t.Length())
		log.Println("The first", i+1, "Titles of Posts:", title)
	})
	/*
	   t := doc.Find(".topics .topic")
	   log.Printf("%+v", t)
	   t = doc.Find(".topics")
	   log.Printf("%+v", t)
	   t = doc.Find(".topic")
	   log.Printf("%+v", t)
	   t = doc.Find("div.topic")
	   log.Printf("div.topic:%+v", t)
	*/
	t := doc.Find("div.topic").Find(".title a")
	log.Printf("div.topic.title a:%+v", t)
	for i := 0; i < t.Length(); i++ {
		d, _ := t.Eq(i).Attr("href")
		title, _ := t.Eq(i).Attr("title")
		fmt.Println(d)
		fmt.Println(title)
	}
}

func CheckError(er error) {
	if er != nil {
		log.Fatal(er)
	}
}
