package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	// Imagine custom logic for generating english greeting
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	// Imagine custom logic for generating spanish greeting
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
