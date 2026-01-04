package main

import (
	"fmt"
)

type SpanishBot struct {
	greeting string
}

func (eb SpanishBot) getGreeting() string {
	return eb.greeting
}

func printGreeting(eb SpanishBot) {
	fmt.Println(eb.getGreeting())
}
