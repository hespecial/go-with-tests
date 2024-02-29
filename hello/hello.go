package main

import (
	"fmt"
	"strings"
)

const helloPrefix = "Hello"

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return strings.Join([]string{helloPrefix, name}, " ")
}

func main() {
	fmt.Println(Hello("world"))
}
