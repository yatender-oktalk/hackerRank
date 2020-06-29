package creational

// BuildProcess for builder interface
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// ManufacturingDirector for builder
type ManufacturingDirector struct {
	builder BuildProcess
}

// Construct for builder
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels().GetVehicle()
}

// SetBuilder for builder
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// VehicleProduct for builder
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// CarBuilder for builder
type CarBuilder struct {
	v VehicleProduct
}

// SetSeats will set seat
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetWheels will set wheels
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetStructure will set structure
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

// GetVehicle for builder
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// BikeBuilder for builder
type BikeBuilder struct {
	v VehicleProduct
}

// SetSeats will set seat
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

// SetWheels will set wheels
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

// SetStructure will set structure
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

// GetVehicle for builder
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
