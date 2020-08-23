package main

import "fmt"

func main() {

	m := map[string]int{
		"Mi": 53,
		"Ma": 55,
		"Di": 28,
	}

	fmt.Println(m)

	// add a new element
	m["Em"] = 30

	for k, v := range m {

		fmt.Println(k, v)

	}

}
