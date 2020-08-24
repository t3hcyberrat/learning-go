package main

import "fmt"

// Hands-on exercise #4
// Create and use an anonymous struct.

func main() {

	as := struct {
		mp map[string]int
		sl []int
	}{
		mp: map[string]int{
			"one":          1,
			"ten":          10,
			"thirty three": 33,
		},
		sl: []int{
			43,
			234,
			22,
		},
	}

	fmt.Println(as)

	for k, v := range as.mp {

		fmt.Println(k, v)

	}

	for i, v := range as.sl {

		fmt.Println(i, v)

	}

}
