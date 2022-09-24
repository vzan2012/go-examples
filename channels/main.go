package main

import (
	"fmt"
	"net/http"
	"time"
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

	// Making an for loop - range of channels (c)
	for l := range c {
		// Passing the link as an argument 
		go func(link string) {
			// Call the Link Checker every after 5 seconds 
			time.Sleep(5 * time.Second)
			linkChecker(link, c)
		}(l)
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
