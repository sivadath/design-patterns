package Builder

import (
	"testing"
)

func TestBuildPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}

	manufacturingComplex.SetBuilder(carBuilder)

	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Error("BuildPattern failed to create car product with 4 wheels. Built wheels count = ", car.Wheels)
	}

	if car.Seat != 5 {
		t.Error("BuildPattern failed to create car product with 5 seats. Built seats count = ", car.Seat)
	}

	if car.Structure != "Car" {
		t.Error("BuildPattern failed to create car product with proper structure. Built structure = ", car.Structure)
	}

	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)

	manufacturingComplex.Construct()

	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Error("BuildPattern failed to create car product with 2 wheels. Built wheels count = ", bike.Wheels)
	}

	if bike.Seat != 2 {
		t.Error("BuildPattern failed to create car product with 2 seats. Built seats count = ", bike.Seat)
	}

	if bike.Structure != "Bike" {
		t.Error("BuildPattern failed to create car product with proper structure. Built structure = ", bike.Structure)
	}
}
