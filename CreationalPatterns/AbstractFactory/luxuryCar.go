package AbstractFactory


type LuxuryCar struct {

}

func (LC *LuxuryCar) NumWheels() int {
	return 4
}

func (LC *LuxuryCar) NumSeats() int {
	return 5
}

func (LC *LuxuryCar) NumDoors() int {
	return 4
}

