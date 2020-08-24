package main

import "fmt"

func main() {

	// anan structs should be used if you only need it in one small place
	// prevents code pollution
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)

}
