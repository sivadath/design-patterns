package AbstractFactory

type FamilyCar struct {

}

func (FC *FamilyCar) NumDoors() int {
	return 5
}

func (FC *FamilyCar) NumWheels() int {
	return 4
}

func (FC *FamilyCar) NumSeats() int {
	return 5
}

