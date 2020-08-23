package main

import "fmt"

// Hands-on exercise #2
// Using a COMPOSITE LITERAL:
// - create a SLICE of TYPE int
// - assign 10 VALUES
// Range over the slice and print the values out.
// Using format printing
// - print out the TYPE of the slice

func main() {

	xi := []int{1, 3, 4, 5, 6, 7, 10, 22, 45, 12, 44, 99}

	for i, v := range xi {

		fmt.Println(i, v)

	}

	fmt.Printf("%T\n", xi)

}
