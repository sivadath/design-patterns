package AbstractFactory

type CruiseBike struct {

}

func (CB *CruiseBike) NumWheels() int {
	return 2
}

func (CB *CruiseBike) NumSeats() int {
	return 2
}

func (CB *CruiseBike) GetBikeType() int {
	return CruiseBikeType
}


