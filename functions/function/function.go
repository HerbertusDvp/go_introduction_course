package function

import "errors"

func Suma(a int, b int) int {
	return a + b
}

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Calc(op Operation, a, b float64) (float64, error) {

	switch op {
	case SUM:
		return a + b, nil
	case SUB:
		return a - b, nil
	case DIV:
		if b == 0 {
			return 0, errors.New("Division zero")
		}
		return a / b, nil
	case MUL:
		return a * b, nil
	}

	return 0, errors.New("Non-existent option")
}
