package Builder

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seat = 2
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

