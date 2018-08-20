package main

import "fmt"

func main() {
	foo(42)
}

func foo(x int) {
	{
		y := 42
		x := 100 + y
		fmt.Printf("Inside this code block x is %d\n", x)
	}

	fmt.Printf("Inside the func level scope x is %d\n", x)
}
