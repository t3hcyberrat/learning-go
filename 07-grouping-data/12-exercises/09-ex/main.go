package main

import "fmt"

// Hands-on exercise #9
// Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

func main() {

	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["noveria"] = []string{"Cold", "Corps", "Hackers"}

	for k, v := range m {
		fmt.Println("Record:", k)

		for i, sv := range v {
			fmt.Printf("\t index: %v \t value: %v \n", i, sv)
		}
	}

}
