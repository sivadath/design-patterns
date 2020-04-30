package AbstractFactory

import (
	"errors"
	"fmt"
)

const (
	SportsBikeType = iota
	CruiseBikeType
)


type Bike interface {
	GetBikeType() int
}

type BikeFactory struct {

}

func (BF *BikeFactory) NewVehicle(bikeType int) (Vehicle, error) {
	switch bikeType {
	case SportsBikeType:
		return new(SportBike), nil
	case CruiseBikeType:
		return new(CruiseBike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Bike of type %d not found",bikeType))
	}
}

