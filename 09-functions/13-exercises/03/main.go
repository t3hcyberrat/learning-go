package main

import "fmt"

// Hands-on exercise #3
// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

func main() {

	defer foo()

	fmt.Println("Printed first")
	fmt.Println("Printed second")

}

func foo() {

	defer func() {

		fmt.Println("Ran after everything else")

	}()

	fmt.Println("foo ran")

}
