package main

import "fmt"

func main() {

	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("done.")

	j := 1
	for {
		if j > 9 {
			break
		}

		fmt.Println(j)
		j++
	}

	fmt.Println("done again.")

}
