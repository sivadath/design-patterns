package AbstractFactory

import (
	"errors"
	"fmt"
)

const (
	CarFactoryType  = iota
	BikeFactoryType
)

type Vehicle interface {
	NumWheels() int
	NumSeats() int
}

type VehicleFactory interface {
	NewVehicle(vehicleType int) (Vehicle, error)
}

func NewVehicleFactory(vehicleFactoryType int) (VehicleFactory, error)  {
	switch vehicleFactoryType {
	case CarFactoryType:
		return  new(CarFactory), nil
	case BikeFactoryType:
		return new(BikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle factory of type %d not found",vehicleFactoryType))
	}
}
