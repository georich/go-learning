package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("This should be deferred")
}

func bar() {
	fmt.Println("This is not deferred")
}
