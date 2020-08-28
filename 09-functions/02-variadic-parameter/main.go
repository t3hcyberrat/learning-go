package main

import "fmt"

func main() {

	add := sum(2, 4, 6, 1, 2, 3)
	fmt.Println(add)

}

func sum(x ...int) int {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {

		sum += v

	}

	return sum

}
