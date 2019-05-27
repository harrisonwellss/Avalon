package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	label = flag.String("label", "img", "label to download")
)

var labelAttrMap = map[string]string{
	"img":    "src",
	"script": "src",
	"a":      "href",
}

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
	doc.Find(*label).Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr(labelAttrMap[*label])
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

func downloadImgs(urls []string, dir string) error {
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return errors.New(resp.Status)
		}

		fullname := filepath.Join(dir, path.Base(u))
		f, err := os.Create(fullname)
		if err != nil {
			return err
		}

		defer f.Close()
		io.Copy(f, resp.Body)
	}
	return nil
}

func maketar(dir string, w io.Writer) error {
	basedir := filepath.Base(dir)

	compress := gzip.NewWriter(w)
	defer compress.Close()

	tr := tar.NewWriter(w)
	defer tr.Close()

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		p, _ := filepath.Rel(dir, path)
		//fmt.Printf("dir:%s, name:%s, p:%s", dir, path, p)

		header.Name = filepath.Join(basedir, p)
		tr.WriteHeader(header)
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		io.Copy(tr, f)
		return nil

	})
	return nil

}

func main() {
	flag.Parse()
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

	tmpdir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tmpdir)
	//defer os.Remove(tmpdir)

	err = downloadImgs(urls, tmpdir)
	if err != nil {
		log.Panic(err)
	}

	f, err := os.Create("img.tar.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	maketar(tmpdir, f)
}

// http://m.xiaohuar.com/
// http://daily.zhihu.com/
