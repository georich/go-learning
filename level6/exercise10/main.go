package main

import "fmt"

func main() {
	foo()
}

func foo() {
	x := 42
	{
		y := 42
		x := 100 + y
		fmt.Printf("Inside this code block x is %d\n", x)
	}

	fmt.Printf("Inside the func level scope x is %d\n", x)
}
