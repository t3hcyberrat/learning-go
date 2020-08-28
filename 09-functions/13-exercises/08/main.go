package main

import "fmt"

// Hands-on exercise #8
// Create a func which returns a func
// assign the returned func to a variable
// call the returned func

func main() {

	a := foo()
	fmt.Println(a())

}

func foo() func() int {
	return func() int {
		return 123
	}
}
