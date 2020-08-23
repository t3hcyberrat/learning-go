package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	// If something doesn't exist in the map, it will return the zero value
	// Here's how to properly check if an item exists in the map
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Barnabas"]; ok {

		fmt.Println("THIS IS THE IF PRINT", v)

	}

}
