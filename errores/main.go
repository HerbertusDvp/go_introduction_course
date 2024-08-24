package main

import (
	"errors"
	"fmt"

	"github.com/herbert/goIntroduction/errores/operator"
)

func main() {
	var err error
	err = errors.New("My new error")
	fmt.Println(err)

	defer func() {
		fmt.Println("In main defer")
		r := recover()
		if r != nil {
			fmt.Println("Recovered in ", r)
		}
	}()

	fmt.Println(err)

	x := 4.0
	y := 0.0
	z := operator.Div(x, y)

	fmt.Println("Result: ", z)

}
