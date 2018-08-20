package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is from an anonymous function")
	}()
}
