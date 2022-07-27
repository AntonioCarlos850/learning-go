package main

import "fmt"

type vehicle struct {
	ports int
	color string
}

type pickup struct {
	vehicle
	four_wheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	car1 := pickup{
		vehicle: vehicle{
			ports: 2,
			color: "red",
		},
		four_wheels: true,
	}

	car2 := sedan{
		vehicle: vehicle{
			ports: 4,
			color: "white",
		},
		luxury: true,
	}

	fmt.Println(car1)
	fmt.Println(car2)

	fmt.Println(car1.color)
	fmt.Println(car2.luxury)
	fmt.Println(car2.vehicle.color)

}
