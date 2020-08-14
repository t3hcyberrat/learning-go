package main

import "fmt"

func main() {

	s := "Hello_world"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// Raw string literal - it includes spaces and all things inside, cannot be done with double quetes
	// rs := `Hello
	// world`
	// fmt.Println(rs)
	// fmt.Printf("%T\n", rs)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("At index position %v we have hex %#x\n", i, v)
	}

}
