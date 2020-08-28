package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(xi...)

	fmt.Println("all numbers:", s)

	e := even(sum, xi...)
	fmt.Println("even numbers:", e)

	o := odd(sum, xi...)
	fmt.Println("odd numbers:", o)

}

func sum(xi ...int) int {

	total := 0

	for _, v := range xi {
		total += v
	}

	return total

}

func even(f func(xi ...int) int, vi ...int) int {

	var yi []int

	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)

}

func odd(f func(xi ...int) int, vi ...int) int {

	var yi []int

	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)

}
