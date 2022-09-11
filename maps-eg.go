package main

import "fmt"

func callColors() {
	// 1. Way to create a map
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"blue":  "#0000ff",
	// 	"green": "#00ff00",
	// 	"grey":  "#cccccc",
	// 	"black": "#000000",
	// }

	// 2. Another way to create a map
	// var colors map[string]string;

	// 3. Using make to create a map
	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	fmt.Println(colors)

	// To Delete a key
	delete(colors, "white")

	fmt.Println(colors)
}
