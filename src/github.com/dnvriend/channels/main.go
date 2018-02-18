package main

import (
	"net/http"
	"fmt"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://github.com",
		"http://tweakers.net",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for link := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(link)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!, cause", err.Error())
		c <- link
	} else {
		fmt.Println(link, "is up!")
		c <- link
	}
}
