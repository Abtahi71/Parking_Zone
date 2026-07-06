package types

type ReservationStatus string

const (
	Active ReservationStatus = "active"
	Completed ReservationStatus = "completed"
	Cancelled ReservationStatus = "cancelled"
)

