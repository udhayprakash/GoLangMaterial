package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Image ...
type Image struct {
	URL      string
	FileName string
}

// Download ...Download
func (image *Image) Download() {
	bytes := GetSource(image.URL, map[string]string{
		"content-type": "image/gif",
	})
	err := ioutil.WriteFile(image.FileName, bytes, 0666)

	HandlerErr(err)
}

// GetSource ...
func GetSource(url string, headers map[string]string) []byte {
	tr := &http.Transport{
		IdleConnTimeout: 30 * time.Second,
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)
	HandlerErr(err)
	for k := range headers {
		req.Header.Set(k, headers[k])
	}

	resp, err := client.Do(req)
	HandlerErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body

}

// HandlerErr handler the err
func HandlerErr(err error) {
	if err != nil {
		panic(err)
	}

}

func main() {
	//url := "https://pixabay.com/images/search/flower/"
	var wg sync.WaitGroup

	var imgs []Image

	url := "https://www.wired.com/2009/08/simpleoutdoorplay/"
	html := string(GetSource(url, map[string]string{
		"User-Agent": "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:80.0) Gecko/20100101 Firefox/80.0",
	}))

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	HandlerErr(err)

	dom.Find("img").Each(
		func(i int, selection *goquery.Selection) {
			imgURL, IsExists := selection.Attr("src")
			if IsExists && strings.HasPrefix(imgURL, "http") {
				strlist := strings.Split(imgURL, "/")

				imgs = append(imgs, Image{URL: imgURL, FileName: strlist[len(strlist)-1]})

			}
		})

	fmt.Println(imgs)
	wg.Add(len(imgs))
	for _, im := range imgs {

		go func(img Image) {
			img.Download()
			wg.Done()
		}(im)

	}
	wg.Wait()

}
