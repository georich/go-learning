package main

import "fmt"

type person struct {
	first    string
	last     string
	flavours []string
}

func main() {
	p1 := person{
		first: "George",
		last:  "Richards",
		flavours: []string{
			"Vanilla",
			"Chocolate",
			"Banana",
		},
	}

	fmt.Println(p1.first, p1.last)
	for _, v := range p1.flavours {
		fmt.Println(v)
	}
}
