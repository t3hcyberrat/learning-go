package main

import "fmt"

func main() {

	switch {

	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case 4 == 4:
		fmt.Println("also true")
	case (7 == 3):
		fmt.Println("not true")
		fallthrough
	case (9 == 3):
		fmt.Println("not true 2")
		fallthrough
	case (15 == 15):
		fmt.Println("true")
	default:
		fmt.Println("this is the default")

	}

	switch "Bond" {

	case "Miss Money":
		fmt.Println("This is Miss Money")
	case "Bond":
		fmt.Println("This is Bond")
	case "Q":
		fmt.Println("This is Q")
	default:
		fmt.Println("This is default")

	}

	switch "Bond" {

	case "Miss Money", "Bond", "Do No":
		fmt.Println("This Miss Money, Bond and Do No")
	case "M":
		fmt.Println("This is M")
	case "Q":
		fmt.Println("This is Q")
	default:
		fmt.Println("This is default")

	}

}
