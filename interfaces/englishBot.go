package main

type EnglishBot struct {
	greeting string
}

func (e EnglishBot) getGreeting() string {
	return e.greeting
}
