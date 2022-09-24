package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.netflix.com",
		"http://bing.com",
	}

	for _, link := range links {
		linkChecker(link)
	}
}

func linkChecker(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down !!!")
		return
	}
	fmt.Println(link, "is up")
}
