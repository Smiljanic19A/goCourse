package main

type SpanishBot struct {
	greeting string
}

func (s SpanishBot) getGreeting() string {
	return s.greeting
}
