package main

import "fmt"

func main() {

	x := []int{1, 4, 55, 33, 232}
	fmt.Println(x)

	// Append to x new values
	x = append(x, 77, 66, 4, 88)
	fmt.Println(x)

	y := []int{123, 5543, 234, 53}
	fmt.Println(y)

	x = append(x, y...)
	fmt.Println(x)

	// Delete from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}
