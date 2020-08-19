package main

import "fmt"

func main() {

	x := 83 / 40
	y := 83 % 40

	fmt.Println(x, y)

	j := 1
	for {
		j++

		if j > 100 {
			break
		}

		if j%2 != 0 {
			continue
		}

		if j%2 == 0 {
			fmt.Println(j)
		}

	}

	fmt.Println("done")

}
