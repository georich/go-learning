package main

import "fmt"

func main() {
	for i := 65; i < 91; i++ {
		fmt.Printf("%v\n", i)
		for c := 1; c < 4; c++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
