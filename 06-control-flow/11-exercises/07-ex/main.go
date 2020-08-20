package main

import "fmt"

// Hands-on exercise #7
// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

func main() {

	x := 2

	if 1 == x {
		fmt.Println("will not print")
	} else if 2 == x {
		fmt.Println("will print")
	} else {
		fmt.Println("will print")
	}

}
