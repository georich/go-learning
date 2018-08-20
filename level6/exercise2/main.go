package main

import "fmt"

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(y []int) int {
	sum := 0
	for _, v := range y {
		sum += v
	}
	return sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := foo(a...)
	fmt.Println(b)

	c := bar(a)
	fmt.Println(c)
}
