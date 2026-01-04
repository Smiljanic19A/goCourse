package main

import (
	"fmt"
)

type EnglishBot struct {
	greeting string
}

func (eb EnglishBot) getGreeting() string {
	return eb.greeting
}

func printGreetingEn(eb EnglishBot) {
	fmt.Println(eb.getGreeting())
}
