package main

import "fmt"

func main() {

	fn := func() int {
		return 42
	}

	x := foo(fn, 24)
	fmt.Println(x)
}

func foo(f func() int, y int) int {
	a := f() + y
	return a
}
