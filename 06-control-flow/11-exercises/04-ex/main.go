package main

import "fmt"

// Hands-on exercise #4
// Create a for loop using this syntax
// for { }
// Have it print out the years you have been alive.

func main() {

	bd := 1990
	for {

		if bd > 2020 {
			break
		}

		fmt.Println(bd)

		bd++

	}

}
