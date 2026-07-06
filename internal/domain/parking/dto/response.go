package dto

import "gotickets/internal/domain/parking/types"

type CreateResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data    CreateData `json:"data"`

}
type CreateData struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Type           types.ParkingType `json:"type"`
	Total_Capacity int    `json:"total_capacity"`
	Price_per_hour int    `json:"price_per_hour"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
type GetResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data    []GetData `json:"data"`
}
type GetSingleResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data    GetData `json:"data"`
}
type GetData struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Type           types.ParkingType `json:"type"`
	Total_Capacity int    `json:"total_capacity"`
	Available_spots int   `json:"available_spots"`
	Price_per_hour int    `json:"price_per_hour"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}