package main

import "fmt"

func main() {

	s1 := foo()
	fmt.Println(s1)

	// x := bar()

	// fmt.Printf("%T \n", x)

	fmt.Println(bar()())

}

func foo() string {

	return "Hello_world"

}

func bar() func() int {

	return func() int {
		return 451
	}

}
