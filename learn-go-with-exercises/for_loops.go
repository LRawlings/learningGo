package main

import "fmt"

func main() {
	// the most basic loop
	// for initialization statement; conditional statement; modified statement {}
	for i := 0; i <10; i++ {
		fmt.Println(i)
	}

	// itrating over an array
	// define array
	v := []int{1,2,3,4,5}

	// loop over array
	for i := 0; i < len(v); i++ {
		fmt.Printf("%d: %d\n",i,v[i])

	}
}