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

	c := make(chan string)

	for _, link := range links {
		go linkChecker(link, c)
	}

	fmt.Println(<-c)
}

func linkChecker(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down !!!")
		c <- "Site might be down !!!"
		return
	}
	fmt.Println(link, "is up")
	c <- "Site can be up !!!"
}
