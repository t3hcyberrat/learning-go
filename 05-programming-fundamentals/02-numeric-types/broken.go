package main

import "fmt"

var v int
var n flovt64

func broken() {

	v = 2.2
	n = 4.44

	fmt.Println(v)
	fmt.Println(n)

	fmt.Printf("%T\n", v)
	fmt.Printf("%T\n", n)

}
