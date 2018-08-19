package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	trucking := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Red",
		},
		fourWheel: true,
	}

	sedanCar := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		luxury: true,
	}

	fmt.Println(trucking)
	fmt.Println(sedanCar)

	fmt.Printf("The truck has %d doors\n", trucking.doors)
	fmt.Printf("The sedan is a %s colour\n", sedanCar.color)
}
