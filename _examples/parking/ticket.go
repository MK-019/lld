package main

import "time"

type ParkingTicket struct {
	entryTime  time.Time
	slotNumber int
	vehicle    *Vehicle
}

type ParkingTicketI interface {
	ParkingTicketGetters
	ParkingTicketSetters
}

func NewEmptyParkingTicket() ParkingTicket {
	return ParkingTicket{}
}

type ParkingTicketGetters interface {
	GetEntryTime() time.Time
	GetSlotNumber() int
	GetVehicle() *Vehicle
}

type ParkingTicketSetters interface {
	SetEntryTime(time.Time)
	SetSlotNumber(int)
	SetVehicle(Vehicle)
}

func (pt *ParkingTicket) GetEntryTime() time.Time {
	return pt.entryTime
}
func (pt *ParkingTicket) SetEntryTime(time_ time.Time) {
	pt.entryTime = time_
}
func (pt *ParkingTicket) GetSlotNumber() int {
	return pt.slotNumber
}
func (pt *ParkingTicket) SetSlotNumber(sn int) {
	pt.slotNumber = sn
}
func (pt *ParkingTicket) GetVehicle() *Vehicle {
	return pt.vehicle
}
func (pt *ParkingTicket) SetVehicle(v Vehicle) {
	pt.vehicle = &v
}
