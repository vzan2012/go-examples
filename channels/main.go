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

	// Making an infinite loop
	for {
		// Call the link checker again 
		go linkChecker(<-c, c)
	}
}

// Link Checker function
func linkChecker(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down !!!")

		// Passing the link to the channel
		c <- link
		return
	}
	fmt.Println(link, "is up")
	
	// Passing the link to the channel
	c <- link
}
