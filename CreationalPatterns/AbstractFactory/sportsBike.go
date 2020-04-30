package AbstractFactory

type SportBike struct {

}

func (SB *SportBike) NumWheels() int {
	return 2
}

func (SB *SportBike) NumSeats() int {
	return 1
}

func (SB *SportBike) GetBikeType() int {
	return SportsBikeType
}

