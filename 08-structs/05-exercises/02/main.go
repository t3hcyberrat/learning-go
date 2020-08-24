package main

import "fmt"

// Hands-on exercise #2
// Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
// Access each value in the map. Print out the values, ranging over the slice.

type person struct {
	first              string
	last               string
	favIceCreamFlavors []string
}

func main() {

	p1 := person{
		first: "test1_first",
		last:  "test1_last",
		favIceCreamFlavors: []string{
			"test1_f1",
			"test1_f2",
			"test1_f3",
		},
	}

	p2 := person{
		first: "test2_first",
		last:  "test2_last",
		favIceCreamFlavors: []string{
			"test2_f1",
			"test2_f2",
			"test2_f3",
		},
	}

	pm := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range pm {

		fmt.Println(v.first)
		fmt.Println(v.last)

		for _, j := range v.favIceCreamFlavors {

			fmt.Println(j)

		}

		fmt.Println("--------------")

	}

}
