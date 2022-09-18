package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Check the args contains value or not
func checkArgs() {
	if len(os.Args) ==1 {
		fmt.Println("You didn't enter the filename")
		os.Exit(1)
	}	
}

type logWriter struct {}

func (logWriter) Write(bs[] byte) (int,error) {
	fmt.Print(string(bs))
	return len(bs),nil
}


func main() {
	checkArgs()

	file, error := os.Open(os.Args[1])
	if(error != nil) {
		log.Fatal(error)
	}

	// Solution 1
	io.Copy(os.Stdout,file)

	// Solution 2
	// lw := logWriter{}
	// io.Copy(lw,file)
}