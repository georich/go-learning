package main

import "fmt"

func main() {
	cat := struct {
		color string
		age   int
		cute  bool
	}{
		color: "Ginger",
		age:   9,
		cute:  true,
	}

	fmt.Println(cat)
}
