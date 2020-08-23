package main

import "fmt"

// Hands-on exercise #3
// Using the code from the previous example, use SLICING to create the following new slices which are then printed:
// - [42 43 44 45 46]
// - [47 48 49 50 51]
// - [44 45 46 47 48]
// - [43 44 45 46 47]

func main() {

	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	a := xi[:5]
	fmt.Println(a)

	b := xi[5:]
	fmt.Println(b)

	c := xi[2:7]
	fmt.Println(c)

	d := xi[1:6]
	fmt.Println(d)

}
