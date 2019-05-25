package main

import "fmt"

const (
	spanish      = "spanish"
	french       = "french"
	englishHello = "Hello"
	spanishHello = "Hola"
	frenchHello  = "Bonjour"
)

func Hello() string {
	return "Hello, world"
}

func HelloYou(name string) string {
	if name == "" {
		name = "stranger"
	}

	return "Hello, " + name
}

func HelloYouLang(name, lang string) string {
	if name == "" {
		name = "stranger"
	}

	return greetingPreifix(lang) + ", " + name
}

func greetingPreifix(lang string) string {
	var hello string

	switch lang {
	case spanish:
		hello = spanishHello
	case french:
		hello = frenchHello
	default:
		hello = englishHello
	}

	return hello
}

func main() {
	fmt.Println(Hello())
}
