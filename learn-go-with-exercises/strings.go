package main

import (
	"fmt"
	"strings"
)

func main() {

	// Declare and then assign
	var car string
	car = "Classic Mini"
	fmt.Printf(car)

	// Or declare and assign at the same time (type is infered)
	var colour = "Red"
	fmt.Printf(colour)

	// New line can be printed using the new line charachter \n
	var animal = "Mouse\nA very small mouse"
	fmt.Printf(animal)

	// Or using Println x times
	fmt.Println(colour)
	fmt.Println(colour)

	// Note: The zero value of type string is a string of zero length, that is an empty string ""


	// Contents of string (bytes) is obtained by standard indexing
	var str = "house"
	fmt.Println(str[0]) // prints 104 (lower case h)
	fmt.Println(str[3]) // prints 115 (lower case s)
	fmt.Println(str[len(str)-1]) // prints 101 (lower case e)

	// Join strings using the package "strings"
	s := strings.Join([]string{"hello", "world"}, ", ")
	fmt.Println(s)

	// Print variables, pass in an integer and a string
	fmt.Printf("%d:%s\n", 2020, "year")

	exercises()
}

func exercises(){
	var favFood string
	var favDrink string

	favFood = "Pizza"
	favDrink = "Cider"

	fmt.Printf("My favourite food is %s\nMy favourite drink is %s", favFood, favDrink)

	name := "Luke"

	fmt.Println("May name is", name)
}