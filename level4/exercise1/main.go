package main

import "fmt"

func main() {
	a := [5]int{11, 13, 24, 42, 100}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)
}
