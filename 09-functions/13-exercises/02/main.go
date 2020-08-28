package main

import "fmt"

// Hands-on exercise #2
// create a func with the identifier foo that
// takes in a variadic parameter of type int
// pass in a value of type []int into your func (unfurl the []int)
// returns the sum of all values of type int passed in
// create a func with the identifier bar that
// takes in a parameter of type []int
// returns the sum of all values of type int passed in

func main() {

	numList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	f := foo(numList...)
	fmt.Println(f)

	b := bar(numList)
	fmt.Println(b)

}

func foo(x ...int) int {

	res := 0
	for _, v := range x {
		res += v
	}

	return res

}

func bar(z []int) int {

	res := 0
	for _, v := range z {
		res += v
	}

	return res

}
