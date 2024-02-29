package main

import (
	"fmt"
	"strings"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
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
	return strings.Join([]string{greetingPrefix(language), name}, ", ")
}

func main() {
	fmt.Println(Hello("Hespecial", french))
}
