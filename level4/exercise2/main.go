package main

import "fmt"

func main() {
	a := []int{11, 15, 19, 21, 29, 37, 50, 78, 99, 100}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)
}
