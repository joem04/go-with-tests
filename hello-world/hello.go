package main

import (
	"fmt"
)

const ( // declared in a block
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {

	if name == "" { // default name to "World"
		name = "World"
	}

	return greetingsPrefix(language) + name

	// prefix := englishHelloPrefix // := declares and assigns, whereas = is just for assigning already declared
}

func greetingsPrefix(language string) (prefix string) {
	// named return value (prefix string), creates return variable upfront and
	// allows for just "return" to be used, it essentially presets the return value
	switch language { // switch case for langauges
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default: // default keyword for switch case
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Napoleon", "French"))
}
