package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak() string
}

func (p *person) speak() string {
	s := fmt.Sprintf("Hello, my name is %s %s", p.first, p.last)
	return s
}

func saySomething(h human) {
	fmt.Println("speak():", h.speak())
}

func main() {
	p1 := person{
		first: "George",
		last:  "Richards",
	}
	saySomething(&p1)
	// saySomething(p1)
}
