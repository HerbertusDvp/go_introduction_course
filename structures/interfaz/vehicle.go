package interfaz

import "fmt"

type Vehicle interface {
	Distance() float64
}

const (
	CarVehicle        = "Car"
	MotorCicleVehicle = "MotorCicle"
	TruckVehicle      = "Truck"
)

func New(v string, time int) (Vehicle, error) {
	switch v {
	case CarVehicle:
		return &Car{Time: time}, nil
	case MotorCicleVehicle:
		return &MotorCicle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	}
	return nil, fmt.Errorf("Vehicle '%s' not exist", v)
}

type Car struct {
	Time int
}

func (c *Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

type MotorCicle struct {
	Time int
}

func (c *MotorCicle) Distance() float64 {
	return 120 * (float64(c.Time) / 60)

}

type Truck struct {
	Time int
}

func (c *Truck) Distance() float64 {
	return 70 * (float64(c.Time) / 60)
}
