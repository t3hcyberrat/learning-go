package main

import "fmt"

func main() {

	fmt.Println("Hello_world")

	// call another function
	foo()

	fmt.Println("Another line")

	// Flow control: iterative
	for i := 0; i < 100; i++ {

		// Flow control: conditional
		if i%2 == 0 {

			fmt.Println("i is", i, "and is an even number")

		}
	}

	bar()

}

func foo() {

	fmt.Println("I'm in foo, hu3hu3")

}

func bar() {

	fmt.Println("...and then we exited")

}

// Control flow
// 1 - sequence
// 2 - loop / iterative
// 3 - conditional
