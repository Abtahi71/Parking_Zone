package types

type ParkingType string

const (
	General ParkingType = "general"
	EV_charging ParkingType = "ev_charging"
	Covered ParkingType = "covered"
)