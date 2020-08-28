package main

import "fmt"

// Hands-on exercise #4
// Create a user defined struct with
// - the identifier “person”
// - the fields:
//   - first
//   - last
//   - age
// - attach a method to type person with
//   - the identifier “speak”
//   - the method should have the person say their name and age
// - create a value of type person
//   - call the method from the value of type person

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {

	fmt.Println("Hello, my name is", p.first, p.last, "and I am", p.age)

}

func main() {

	p1 := person{
		first: "Tamagochi",
		last:  "Suki",
		age:   30,
	}

	p1.speak()

}
