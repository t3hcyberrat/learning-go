package main

import "fmt"

func main() {

	//
	x := []int{1, 2, 4, 5, 100, 32, 42, 221}

	fmt.Println(len(x))
	fmt.Println(x)

	for i, v := range x {

		fmt.Println(i, v)

	}

}
