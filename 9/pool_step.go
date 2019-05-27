package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func printUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	fmt.Println(url, resp.StatusCode)
}

func work(ch chan string, wg *sync.WaitGroup) {
	for url := range ch {
		printUrl(url)
	}

	wg.Done()

}

func main() {

	ch := make(chan string)

	wg := new(sync.WaitGroup)
	wg.Add(3)

	for i := 0; i < 3; i++ {
		//wg.Add(1)
		go work(ch, wg)
	}

	for i := 0; i < 5; i++ {
		url := "http://www.baidu.com"
		ch <- url
	}
	close(ch)

	wg.Wait()

	time.Sleep(3 * time.Second)
}
