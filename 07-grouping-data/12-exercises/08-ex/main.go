package main

import "fmt"

// Hands-on exercise #8
// Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things.
// Store three records in your map. Print out all of the values, along with their index position in the slice.

// `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
// `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
// `no_dr`, `Being evil`, `Ice cream`, `Sunsets`

func main() {

	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println("key: ", k)

		for i, sval := range v {
			fmt.Printf("\t index position: %v \t value: %v \n", i, sval)
		}
	}

}
