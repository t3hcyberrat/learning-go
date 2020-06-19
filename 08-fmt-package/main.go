package main

import "fmt"

var y = 911

func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#X\n", y)

	y = 100

	fmt.Printf("%#X", y)
	fmt.Printf("%#X\t%b\t%x\n", y, y, y)

}
