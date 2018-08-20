package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hello I am %s %s and I am %d years old", p.first, p.last, p.age)
}

func main() {
	g := person{
		first: "George",
		last:  "Richards",
		age:   24,
	}

	g.speak()
}
