package main

import "fmt"

func main() {
	today := "Friday"
	switch today {
	case "Saturday":
		fmt.Println("Let's relax for the weekend")
	case "Sunday":
		fmt.Println("Lets relax for the weekend")
	case "Friday":
		fmt.Println("Friday night! Let's party!")
	default:
		fmt.Println("Back to work...")

	}
}
