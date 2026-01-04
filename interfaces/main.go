package main

import "fmt"

func main() {
	spanish := SpanishBot{
		greeting: "Hola essey",
	}
	english := EnglishBot{
		greeting: "Hey fren",
	}

	printGreeting(english)
	printGreeting(spanish)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
