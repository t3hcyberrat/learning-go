package main

import "fmt"

// Hands-on exercise #5
// Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.

func main() {

	for i := 1; i <= 100; i++ {

		if i%4 == 0 {

			fmt.Println(i)

		}

	}

}
