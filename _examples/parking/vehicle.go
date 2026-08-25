package main

type Vehicle struct {
	licensePlate string
	color        string
}

func NewVehicle(plate, color string) Vehicle {
	return Vehicle{
		licensePlate: plate,
		color:        color,
	}
}

type VehicleInterface interface {
	VehicleGetter
	VehicleSetter
}

type VehicleGetter interface {
	GetLicensePlate() string
	GetColor() string
}

type VehicleSetter interface {
	SetLicensePlate(string)
	SetColor(string)
}

func (v *Vehicle) GetLicensePlate() string {
	return v.licensePlate
}
func (v *Vehicle) GetColor() string {
	return v.color
}

func (v *Vehicle) SetColor(color string) {
	v.color = color
}
func (v *Vehicle) SetLicensePlate(li string) {
	v.licensePlate = li
}
