package main

import "fmt"

// Hands-on exercise #4
// Follow these steps:
// - start with this slice
//   - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// - append to that slice this value
//   - 52
// - print out the slice
// - in ONE STATEMENT append to that slice these values
//   - 53
//   - 54
//   - 55
// - print out the slice
// - append to the slice this slice
//   - y := []int{56, 57, 58, 59, 60}
// - print out the slice

func main() {

	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	xi = append(xi, 52)
	fmt.Println(xi)

	xi = append(xi, 53, 54, 55)
	fmt.Println(xi)

	y := []int{56, 57, 58, 59, 60}
	xi = append(xi, y...)
	fmt.Println(xi)

}
