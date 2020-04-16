package Builder


type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seat = 5
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

