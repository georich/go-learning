package main

import "fmt"

func main() {
	x := returnFunc()
	fmt.Println(x())
}

func returnFunc() func() int {
	return func() int {
		return 42
	}
}
