package main

import "fmt"

func main() {
	a := 3 == 2
	b := 3 <= 2
	c := 3 >= 2
	d := 3 != 2
	e := 3 > 2
	f := 3 > 2
	fmt.Printf("The values are: %v %v %v %v %v %v", a, b, c, d, e, f)
}
