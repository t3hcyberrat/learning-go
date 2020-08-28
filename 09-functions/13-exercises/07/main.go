package main

import "fmt"

// Hands-on exercise #7
// Assign a func to a variable, then call that func

func main() {

	z := func(a int, b int) int {
		return a + b
	}

	fmt.Println(z(1, 3))

	for i := 0; i < 100; i++ {
		fmt.Printf("%v\t", i)
	}

}
