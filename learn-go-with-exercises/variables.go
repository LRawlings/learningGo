package main

import (
	"fmt"
	"time"
	"strconv"
)

// When a variable is declared by a var, the system automatically gives it a zero value of that type
// Global variables
var (
	a int //=0
	b bool //=False
	s string //=""
	apple int //=0
	bee bool //=False
	apple2 string //=""

	f float64 //=0.0
	p *int //=nil    pointer
)

func main() {
	// multiple variables assigned on same row (parallel assignments) (:= declares them at the same time as assigning)
	// to just assign previously declared variables, it would be h, i, j = 5, 7, "abc"
	h, i, j := 5, 7, "abc"
	fmt.Println(i, j)

	fmt.Println(h)     // prints 5
    {
        fmt.Println(h) // prints 5
        h := 2
        fmt.Println(h) // prints 2
    }
	fmt.Println(h)     // prints 5
	

	// Strings are ouytputted using the 'special' symbol %s
	str1 := "Hello"
    str2 := "World"
    
	fmt.Printf("%s %s\n",str1,str2)
	

	// Integers are outputted using the 'decimal' symbol %d (base 10)
	// but can also be outputed in other ways like binary (base 2) %b, octal (base 8) %o, or hexadecimal (base 16) %x
    int1 := 1
    int2 := 2
    fmt.Printf("int1 is %d\n", int1)
    fmt.Printf("int2 is %d\n", int2)


	// Float: floating point number i.e. Pi 1.314
	// output using %f
	banana := 3.0
	bread := 2.0
	price := banana + bread

	fmt.Printf("")
	fmt.Printf("Price:     %f\n",price)
	vat := price * 0.15
	fmt.Printf("VAT:     %f\n",vat)
	total := vat + price
	// limit the number of decimal places using %.2f for 2 decimal places
	fmt.Printf("Total:    %.2f\n",total)
	fmt.Printf("")


	// Exchange variables
	// a has been exchanged with b
	a := 1
	b := 2
	a,b = b,a
	fmt.Println(a)
	fmt.Println(b)


	// blank identifier _
	_, b = 5, 7
	// 5 is abandoned


	exercise()
}

func exercise() {
	dob := "05/12/1990"
	var year int
	var age int

	dobDay, _ := strconv.Atoi(dob[0:2])
	dobMonth, _ := strconv.Atoi(dob[3:5])
	dobYear, _ := strconv.Atoi(dob[6:10])

	curDay := time.Now().Day()
	curMonth := int(time.Now().Month())
	curYear := time.Now().Year()

	fmt.Println(dob, year, age, dobDay, dobMonth, dobYear, curDay, curMonth, curYear)

	if dobMonth < curMonth && dobDay < curDay {
		age = curYear - dobYear -1
	} else {
		age = curYear - dobYear
	}

	fmt.Printf("Todays date is %s\n Someone with a date of birth of %s will be %d years old.", time.Now(), dob, age)
}