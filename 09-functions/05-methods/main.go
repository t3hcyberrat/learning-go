package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return) { code }
func (s secretAgent) speak() {

	fmt.Println("I am,", s.first, s.last)

}

func main() {

	p1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	fmt.Println(p1)
	p1.speak()

}
