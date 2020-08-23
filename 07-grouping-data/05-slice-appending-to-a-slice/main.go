package main

import "fmt"

func main() {

	// Declare a var with a slice of ints
	x := []int{1, 4, 55, 33, 232}
	fmt.Println(x)

	// Append to x new values
	x = append(x, 77, 66, 4, 88)
	fmt.Println(x)

	y := []int{123, 5543, 234, 53}
	fmt.Println(y)

	x = append(x, y...)
	fmt.Println(x)

}
