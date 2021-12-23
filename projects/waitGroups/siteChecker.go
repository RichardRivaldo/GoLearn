package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func saveWebsite(url string, waitGroup *sync.WaitGroup) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "is down!")
	} else {
		defer response.Body.Close()

		if response.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(response.Body)
			fileName := strings.Split(url, "//")[1] + ".txt"

			err = ioutil.WriteFile(fileName, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Successfully saved", url)
			}
		}
	}

	waitGroup.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	urls := []string{"https://golang.org", "https://www.google.com"}

	waitGroup.Add(len(urls))

	for _, url := range urls {
		go saveWebsite(url, &waitGroup)
	}

	waitGroup.Wait()
}
