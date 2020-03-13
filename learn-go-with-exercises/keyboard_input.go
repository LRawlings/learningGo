package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
)

func main() {
	// create a new reader from stdin
    reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your city: ")
	// read line from console
	// Passing '\n' as the delim argument stops reading the string when a newline is encountered,
	// like a user pressing enter.
	// seconf return value is an error value (we are ignoring it here using _)
    city, _ := reader.ReadString('\n')
    fmt.Print("You live in " + city)
}

func exercise(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is: ", name)

	fmt.Print("What is your favorate number? ")
	strNumber, _ := reader.ReadString('\n')

	// String conversions. Atoi (string to int) and Itoa (int to string). Second return value is an error
	number, err := strconv.Atoi(strNumber)

	// Error checking
	if err != nil {
		fmt.Println("conversion error:", strNumber)
	}

	if number > 1 && number < 10 {
		fmt.Println("The number you entered was: ", number)
	} else {
		fmt.Println("The number you entered was not between 1 and 10!")
	}
}