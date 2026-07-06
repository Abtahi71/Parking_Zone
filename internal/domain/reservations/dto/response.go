package dto

import (
	reservationTypes "gotickets/internal/domain/reservations/types"
	ParkingType "gotickets/internal/domain/parking/types"
)
type ReserveRes struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data    ReserveData `json:"data"`

}

type ReserveData struct {
	Id            uint   `json:"id"`
	User_id       uint    `json:"user_id"`
	Zone_id       uint    `json:"zone_id"`
	License_plate string `json:"license_plate"`
	Status        reservationTypes.ReservationStatus `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}


type GetReserveRes struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	Data    []GetReserveData `json:"data"`
}
type GetReserveData struct{
   Id             uint   `json:"id"`
   License_plate  string `json:"license_plate"`
   Status         reservationTypes.ReservationStatus `json:"status"`
   Zone           ZoneData `json:"zone"`
   CreatedAt      string `json:"created_at"`
   UpdatedAt      string `json:"updated_at"`
}

type ZoneData struct{
	Id uint `json:"id"`
	Name string `json:"name"`
	Type ParkingType.ParkingType `json:"type"`
}

type CancelResponse struct{
	Success bool `json:"success"`
	Message string `json:"message"`
}