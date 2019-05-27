package main

import (
	"fmt"
	_ "github.com/PuerkitoBio/goquery"
	"github.com/baidubce/bce-sdk-go/util/log"
	"net/url"
	"os"
)

func main() {
	s := os.Args[1]
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.User)
	fmt.Println(u.Fragment)

}
