package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Links Array
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.netflix.com",
		"http://bing.com",
		"http://linkedin.com",
		"http://indeed.com",
	}

	// Creating a channel of type string
	c := make(chan string)

	// Looping the link
	for _, link := range links {
		go linkChecker(link, c)
	}

	// Looping the channel messages
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

// Link Checker function
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
