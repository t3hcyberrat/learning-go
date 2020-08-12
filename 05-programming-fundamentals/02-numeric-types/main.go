package main

import "fmt"

var a int
var b float64

func main() {

	a = 22
	b = 4.44

	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	x := 12
	y := 42.222

	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

}
