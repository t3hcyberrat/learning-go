package main

import "fmt"

// Hands-on exercise #6
// Build and use an anonymous func

func main() {

	func(a int, b int) {
		fmt.Println(a + b)
	}(1, 2)

}
