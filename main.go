package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const dutch = "Dutch"

const spanishHello = "Hola"
const englishHello = "Hello"
const frenchHello = "Bonjour"
const dutchHello = "Hallo"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + ", " + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		return spanishHello
	case french:
		return frenchHello
	case dutch:
		return dutchHello
	default:
		return englishHello
	}
}

func main() {
	fmt.Println(Hello("Moustafa", ""))
}
