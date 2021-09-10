package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%+v", err)
		return ""
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	return ctype
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	channel := make(chan string)

	for _, url := range urls {
		go (func(url string) {
			ctype := contentType(url)
			channel <- ctype
		})(url)
	}

	for range urls {
		ctype := <-channel
		fmt.Println(ctype)
	}
}
