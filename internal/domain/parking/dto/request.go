package dto

import "gotickets/internal/domain/parking/types"

type CreateParking struct {
	Name           string            `json:"name" validate:"required"`
	Type           types.ParkingType `json:"type" validate:"required"`
	Total_Capacity int               `json:"total_capacity" validate:"required"`
	Price_per_hour int               `json:"price_per_hour" validate:"required"`
}
