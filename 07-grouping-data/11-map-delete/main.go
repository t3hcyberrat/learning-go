package main

import "fmt"

func main() {

	m := map[string]int{
		"Mi": 53,
		"Ma": 55,
		"Di": 28,
		"Em": 30,
	}

	fmt.Println(m)

	delete(m, "Em")
	fmt.Println(m)

	// Notice no error was thrown
	delete(m, "Ian Flemming")
	fmt.Println(m)

	// Use the comma ok idiom to check if things we want to delete actually exist
	if v, ok := m["Ian Flemming"]; ok {

		fmt.Println("Deleting key: ", v)
		delete(m, "Ian Flemming")

	} else {

		fmt.Println("Does not exist")

	}

}
