package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func saveWebsite(url string, channel chan string) {
	response, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is down! Error:%s", url, err)
		channel <- s
	} else {
		defer response.Body.Close()

		if response.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(response.Body)
			fileName := strings.Split(url, "//")[1] + ".txt"

			err = ioutil.WriteFile(fileName, bodyBytes, 0664)
			if err != nil {
				// log.Fatal(err)
				channel <- "Error writing the file"
			} else {
				s := fmt.Sprintf("Successfully saved %s", url)
				channel <- s
			}
		}
	}
}

func main() {
	channel := make(chan string)
	urls := []string{"https://golang.org", "https://www.google.com"}

	for _, url := range urls {
		go saveWebsite(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-channel)
	}
}
