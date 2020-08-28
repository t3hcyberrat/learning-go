package main

import "fmt"

func main() {

	foo()

	func() {
		fmt.Println("anon func ran")
	}()

	func(x int) {
		fmt.Println("anon func ran", x)
	}(42)

}

func foo() {

	fmt.Println("foo ran")

}
