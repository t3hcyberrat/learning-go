package main

import "fmt"

// Hands-on exercise #7
// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."

// Range over the records, then range over the data in each record.

func main() {

	j := []string{"James", "Bond", "Shaken, not stirred"}
	m := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xxs := [][]string{j, m}

	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)

		for j, val := range xs {
			fmt.Printf("\t index position: \t %v \t value: \t %v \n", j, val)
		}
	}

}
