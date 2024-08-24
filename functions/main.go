package main

import (
	"encoding/json"
	"fmt"

	"github.com/herbert/goIntroduction/structures/structure"
)

func main() {

	maiEstruct := structure.Product{
		Name:  "Heladera marca sarasa",
		Price: 40000,
		Type: structure.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tag:   []string{"heladera", "sarasa", "freezer", "Fridge"},
		Count: 5,
	}

	v, err := json.Marshal(maiEstruct) //Cnvierte a formato jason

	fmt.Println("Datos: ", string(v))
	fmt.Println("Error:", err)
	fmt.Println("Total:", maiEstruct.TotalPrice())
	fmt.Println()
	maiEstruct.SetName("Fedelobo")
	fmt.Println(maiEstruct)

	v, _ = json.Marshal(maiEstruct) //Cnvierte a formato jason
	fmt.Println("Datos: ", string(v))

}
