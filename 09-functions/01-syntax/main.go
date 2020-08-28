package main

import "fmt"

func main() {

	foo()
	bar("Ugly")

	s1 := woo("Moneypenny")
	fmt.Println(s1)

	x, y := mouse("Ian", "Flemming")
	fmt.Println(x)
	fmt.Println(y)

}

// func (r receiver) identifier(parameters) (return(s)) { code }

func foo() {

	fmt.Println("Hello from foo")

}

// Everything in Go is PASS BY VALUE
func bar(s string) {

	fmt.Println("Hello,", s)

}

func woo(s string) string {

	return fmt.Sprint("Hello from woo, ", s)

}

func mouse(fn string, ln string) (string, bool) {

	s := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false

	return s, b

}
