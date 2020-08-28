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

func (p person) speak() {

	fmt.Println("I am,", p.first, p.last, "-- the person speak")

}

func bar(h human) {

	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first)
	}

}

type human interface {
	speak()
}

func main() {

	p1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p2 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(p1)
	p1.speak()

	fmt.Println(p2)
	p2.speak()

	bar(p1)
	bar(p2)

}
