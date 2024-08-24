package main

import (
	"fmt"

	"github.com/herbert/goIntroduction/structures/interfaz"
)

func main() {
	fmt.Println("Manejo de interfaces")
	fmt.Println()

	auto := interfaz.Car{Time: 120}
	fmt.Println(auto.Distance())

	medios := []string{"Car", "MotorCicle", "Truck"}

	for _, v := range medios {
		fmt.Printf("Vehicle: %s\n", v)
		transporte, err := interfaz.New(v, 400)

		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println()
			continue
		}
		distance := transporte.Distance()
		fmt.Printf("Distance: %.2f\n", distance)
		fmt.Println()
	}

}
