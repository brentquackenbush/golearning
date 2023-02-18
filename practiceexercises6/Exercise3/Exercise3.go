package main

import "fmt"

/*
● Create a user defined struct with
○ the identifier “person”
○ the fields:
■ first ■ last ■ age
● attach a method to type person with
○ the identifier “speak”
○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person
*/

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "Brent",
		last:  "Quackenbush",
		age:   25,
	}
	p.speak()
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am age,", p.age)
}
