package main

import "fmt"

// Hands-on exercise #9
// A “callback” is when we pass a func into a func as an argument. For this exercise,
// pass a func into a func as an argument

func main() {

	fmt.Println(foo(bar))

}

func bar() int {
	return 42
}

func foo(x func() int) int {

	f := x()

	return f

}
