package main

import "fmt"

// Hands-on exercise #3
// Create a for loop using this syntax
// for condition { }
// Have it print out the years you have been alive.

func main() {

	bd := 1990
	for bd < 2020 {
		fmt.Println(bd)
		bd++
	}

}
