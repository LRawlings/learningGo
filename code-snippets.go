package main

import "fmt"

func main() {
	// for loops
	// i = index, v = copy of the element

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}


	pow2 := make([]int, 10)

	// If you only want the index, you can omit the second variable
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // n << x is "n times 2, x times". And y >> z is "y divided by 2, z times"
	}

	// You can skip the index or value by assigning to _
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	for i, _ := range pow2 {
		fmt.Printf("%d\n", i)
	}



}