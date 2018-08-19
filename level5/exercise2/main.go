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

	m := map[string]person{
		p1.last: p1,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.flavours {
			fmt.Println(i, val)
		}
	}
}
