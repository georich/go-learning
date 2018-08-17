package main

import "fmt"

func main() {
	today := "Monday"
	if today == "Friday" {
		fmt.Println("It's Friday! Let's party!")
	} else {
		fmt.Println("Better get back to work...") // Ignoring weekend
	}
}
