package main

func main() {
	spanish := SpanishBot{
		greeting: "Hola essey",
	}
	english := EnglishBot{
		greeting: "Hey fren",
	}

	printGreeting(spanish)
	printGreetingEn(english)
}
