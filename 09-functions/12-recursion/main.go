package main

import "fmt"

func main() {

	fmt.Println(4 * 3 * 2 * 1)

	f := factorial(4)
	fmt.Println(f)

	fl := factorialLoop(4)
	fmt.Println(fl)

}

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)

}

func factorialLoop(n int) int {

	result := 1

	for ; n > 0; n-- {

		result *= n

	}

	return result

}
