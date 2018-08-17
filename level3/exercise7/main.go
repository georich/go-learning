package main

import "fmt"

func main() {
	today := "Sunday"
	if today == "Friday" {
		fmt.Println("Woo! Let's party!")
	} else if today == "Saturday" || today == "Sunday" {
		fmt.Println("Let's relax for the weekend")
	} else {
		fmt.Println("Better get back to work...")
	}
}
