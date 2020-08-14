package main

import "fmt"

// Hands-on exercise #4
// Write a program that
// - assigns an int to a variable
// - prints that int in decimal, binary, and hex
// - shifts the bits of that int over 1 position to the left, and assigns that to a variable
// - prints that variable in decimal, binary, and hex

func main() {

	//
	x := 2
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	//
	bsl := x << 1
	fmt.Printf("%d\t%b\t%#x\n", bsl, bsl, bsl)

}
