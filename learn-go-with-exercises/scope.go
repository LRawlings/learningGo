package main

import "fmt"

// global variable. can be used by multiple functions
var y = 99

func example() {
	fmt.Println(y)
}

func main() {
	// local variable. Can only be used inside function main()
    x := 7
	fmt.Println(x)
	y := 20
	// local variable take priority over global variables
	fmt.Println(y)
	example()

	fmt.Println("x + y = ", sum(x,y))
}

// function parametes (a and b here) are used used a local variables within the function
func sum(a, b int) int {
	return a + b
}