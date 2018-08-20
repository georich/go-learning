package main

import "fmt"

func main() {
	fn := func(s string) {
		fmt.Println(s)
	}

	fn("Hello!")
}
