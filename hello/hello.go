package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "Hello"
const spanish = "Spanish"
const spanishHelloPrefix = "Hola"
const french = "French"
const frenchHelloPrefix = "Bonjour"

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return strings.Join([]string{greetingPrefix(language), name}, " ")
}

func main() {
	fmt.Println(Hello("world", ""))
}
