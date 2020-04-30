package AbstractFactory

import (
	"errors"
	"fmt"
)

const (
	LuxuryCarType = iota
	FamilyCarType
)

type Car interface {
	NumDoors() int
}

type CarFactory struct {

}

func (cf *CarFactory) NewVehicle(carType int) (Vehicle, error) {
	switch carType {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Car type %d not found",carType))
	}
}

