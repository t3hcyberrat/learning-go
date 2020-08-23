package main

import "fmt"

func main() {

	x := []int{1, 2, 44, 23, 32, 42}

	fmt.Println(x)
	fmt.Println(x[1])   // access by index
	fmt.Println(x[1:])  // sliced slice
	fmt.Println(x[1:3]) // sliced slice

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

}
