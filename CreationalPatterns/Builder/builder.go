package Builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetStructure() BuildProcess
	SetSeats() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels	int
	Seat 	int
	Structure string
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (md *ManufacturingDirector) Construct() {
	md.builder.SetSeats().SetStructure().SetWheels()
}

func (md *ManufacturingDirector) SetBuilder(builder BuildProcess) {
	md.builder = builder
}

