package main

import (
	"fmt"
	"strings"
)

func Hello(name string) string {
	return strings.Join([]string{"Hello", name}, " ")
}

func main() {
	fmt.Println(Hello("world"))
}
