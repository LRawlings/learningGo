package main

import "fmt"

func main() {
	// Declare an array with length 5 and type int
	// For type int all elements are automatically initialised to the default value of 0
	var arr1 = new([5]int)
	fmt.Println(arr1[4]) // prints 0

	// Declare and assign at the same time
	var arrAge = [5]int{18, 12, 32, 53, 22}
	fmt.Println(arrAge)

	// array of non-fixed length. Length is determined by number of elements specified at initialisation time
	var arrLazy = [...]int{5,6,7,8,1,5,7,9,7,8,4}
	fmt.Println(arrLazy)

	// array with key values
	var arrKeyVakue = [5]string{3: "John", 4: "James"}

	fmt.Println(arrKeyVakue[0]) //returns 0
	fmt.Println(arrKeyVakue[3]) //returns "John"
}