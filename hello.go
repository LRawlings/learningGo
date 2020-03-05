package main

import "fmt"

const englishHeloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spannish" {
		return "Hola, " + name
	}

	return englishHeloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
