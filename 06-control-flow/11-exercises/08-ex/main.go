package main

import "fmt"

// Hands-on exercise #8
// Create a program that uses a switch statement with no switch expression specified.

func main() {

	x := "test"
	switch {
	case x == "test":
		fmt.Println("this is a test and will print")
	case x == "not test":
		fmt.Println("this is not a test and will not print")
	}

}
