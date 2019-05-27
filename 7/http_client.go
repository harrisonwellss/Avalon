package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func fetch(url string) ([]string, error) {
	var urls []string
	//url := "http://daily.zhihu.com/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	//io.Copy(os.Stdout, resp.Body)

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")
		if ok {
			urls = append(urls, link)
			//fmt.Println(link)
		}
	})
	return urls, nil
}

func cleanUrls(u string, urls []string) []string {
	var resUrl []string

	uri, err := url.Parse(u)
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range urls {
		if strings.HasPrefix(u, "//") {
			u = uri.Scheme + "://" + uri.Host + u
			resUrl = append(resUrl, u)

		} else if strings.HasPrefix(u, "/") {
			u = uri.Scheme + "://" + uri.Host + u
			resUrl = append(resUrl, u)
		} else {
			resUrl = append(resUrl, u)
		}
	}

	return resUrl
}

func main() {

	url := os.Args[1]

	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	//for _, u := range urls{
	//	fmt.Println(u)
	//}

	resUrls := cleanUrls(url, urls)

	for _, u := range resUrls {
		fmt.Println(u)
	}
}

// http://m.xiaohuar.com/
// http://daily.zhihu.com/
