package main

import (
	"errors"
	"time"
)

type ParkingLot struct {
	Capacity     int
	Available    int
	ParkingSlots []ParkingSlot
}

// CreateParkingLot initializes a new parking lot with the given capacity.
func CreateParkingLot(capacity int) *ParkingLot {
	parkingLot := &ParkingLot{
		Capacity:     capacity,
		Available:    capacity,
		ParkingSlots: make([]ParkingSlot, capacity),
	}
	for i := 0; i < capacity; i++ {
		parkingLot.ParkingSlots[i].ID = i + 1
	}
	return parkingLot
}

// ParkVehicle parks a vehicle in the parking lot, if space is available.
// It returns a parking ticket on success or an error if the lot is full.
func (pl *ParkingLot) ParkVehicle(vehicle Vehicle) (ParkingTicket, error) {
	if pl.Available == 0 {
		return ParkingTicket{}, errors.New("parking lot is full")
	}

	slotNumber, err := pl.findAvailableSlot()
	if err != nil {
		return ParkingTicket{}, err
	}

	pl.ParkingSlots[slotNumber-1].Occupied = true
	pl.ParkingSlots[slotNumber-1].Vehicle = vehicle
	pl.Available--

	ticket := NewEmptyParkingTicket()
	ticket.SetEntryTime(time.Now())
	ticket.SetSlotNumber(slotNumber)
	ticket.SetVehicle(vehicle)

	return ticket, nil
}

// findAvailableSlot finds the first available parking slot and returns its number.
// It returns an error if no available slots are found.
func (pl *ParkingLot) findAvailableSlot() (int, error) {
	for i, slot := range pl.ParkingSlots {
		if !slot.Occupied {
			return i + 1, nil
		}
	}
	return 0, errors.New("no available slots")
}

// ExitParkingLot allows a vehicle to exit the parking lot using a parking ticket.
// It marks the slot as vacant and increments the available slots count.
// Returns an error if the ticket is invalid.
func (pl *ParkingLot) ExitParkingLot(ticket ParkingTicket) error {
	if ticket.GetSlotNumber() <= 0 || ticket.GetSlotNumber() > pl.Capacity {
		return errors.New("invalid slot number")
	}

	slot := &pl.ParkingSlots[ticket.GetSlotNumber()-1]
	if !slot.Occupied || slot.Vehicle.GetLicensePlate() != ticket.GetVehicle().GetLicensePlate() {
		return errors.New("invalid parking ticket")
	}

	slot.Occupied = false
	slot.Vehicle = Vehicle{}
	pl.Available++

	return nil
}
