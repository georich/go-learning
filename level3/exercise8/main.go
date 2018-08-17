package main

import "fmt"

func main() {
	switch {
	case (2 == 3):
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print")
	case (2 == 2):
		fmt.Println("This should print!")
	case (2 == 5):
		fmt.Println("This should not print")
	}
}
