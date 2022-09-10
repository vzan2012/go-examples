package main

import "fmt"

func experiment() {
	mySlice := []string{"A", "B", "C", "D"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Word"
}
