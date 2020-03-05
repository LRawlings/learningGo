package main

import "fmt"

const englishHeloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHeloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
