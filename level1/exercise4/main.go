package main

import "fmt"

type pineapple int

var x pineapple

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
