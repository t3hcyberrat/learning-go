package main

import "fmt"

// Hands-on exercise #1
// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
// - first name
// - last name
// - favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

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

	fmt.Println(p1.first, p1.last)

	for _, v := range p1.favIceCreamFlavors {

		fmt.Printf("\t Favourite ice cream: %v \n", v)

	}

	fmt.Println(p2.first, p2.last)

	for _, v := range p2.favIceCreamFlavors {

		fmt.Printf("\t Favourite ice cream: %v \n", v)

	}

}
