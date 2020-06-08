package main

import "fmt"

// GO is a STATIC programming language
// NOT a DYNAMIC programming language
// therefore a VARIABLE is DECLARED to hold a VALUE of a certain TYPE only

var y = 46

// DECLARED that the VARIABLE with IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "Shaken, not stirred"
var z = "Shaken, not stirred"

var a = `James said, "Shaken, not stirred"`

func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// z = 46
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)

}
