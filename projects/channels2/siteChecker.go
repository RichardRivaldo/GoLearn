package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkWebsite(url string, channel chan string) {
	response, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is down! Error:%s", url, err)
		fmt.Println(s)
		channel <- url
	} else {
		defer response.Body.Close()

		if response.StatusCode == 200 {
			s := fmt.Sprintf("%s is up!", url)
			fmt.Println(s)
			channel <- url
		}
	}
}

func main() {
	channel := make(chan string)
	urls := []string{"https://golang.org", "https://www.google.com"}

	for _, url := range urls {
		go checkWebsite(url, channel)
	}

	// for {
	// 	go checkWebsite(<-channel, channel)
	// 	time.Sleep(time.Second * 2)
	// }

	// for url := range channel {
	// 	time.Sleep(time.Second * 2)
	// 	go checkWebsite(url, channel)
	// }

	for url := range channel {
		go func(u string) {
			time.Sleep(time.Second * 2)
			checkWebsite(u, channel)
		}(url)
	}
}
