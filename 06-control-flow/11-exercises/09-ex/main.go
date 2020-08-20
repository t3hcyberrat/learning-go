package main

import "fmt"

// Hands-on exercise #9
// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

func main() {

	favSport := "Biking"
	switch favSport {

	case "Biking":
		fmt.Println("Biking")
	case "Football":
		fmt.Println("Hu3hu3")

	}

}
