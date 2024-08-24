package operator

func Div(x, y float64) float64 {

	defer func() {

		if y == 0 {
			panic("Erro by zero")
		}
	}()

	return x / y

}
