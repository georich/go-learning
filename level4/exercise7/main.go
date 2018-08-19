package main

import "fmt"

func main() {
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Hello, James"}
	ab := [][]string{a, b}
	for i, v := range ab {
		fmt.Println(i, v)
		for index, value := range v {
			fmt.Println(index, value)
		}
	}
}
