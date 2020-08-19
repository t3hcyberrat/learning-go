package main

import "fmt"

func main() {

	i := 33
	for i < 123 {
		fmt.Println(i)
		fmt.Printf("%#U\n", i)
		i++
	}

}
