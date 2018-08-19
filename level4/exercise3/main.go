package main

import "fmt"

func main() {
	a := []int{11, 15, 19, 21, 29, 37, 50, 78, 99, 100}
	fmt.Println(a[2:7])
	fmt.Println(a[2:])
	fmt.Println(a[:7])
}
