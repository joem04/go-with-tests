package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {

	if name == "" { // default name to "World"
		name = "World"
	}

	prefix := englishHelloPrefix // := declares and assigns, whereas = is just for assigning already declared

	switch language { // switch case for langauges
	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Napoleon", "French"))
}
