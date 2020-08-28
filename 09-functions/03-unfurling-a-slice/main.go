package main

import "fmt"

func main() {

	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}

	add := sum(xi...)

	fmt.Println(add)

}

func sum(x ...int) int {

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum

}
