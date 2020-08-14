package main

import "fmt"

// Hands-on exercise #6
// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

const (
	y1 = 2020 + iota
	y2 = 2020 + iota
	y3 = 2020 + iota
	y4 = 2020 + iota
)

func main() {

	fmt.Println(y1, y2, y3, y4)

}
