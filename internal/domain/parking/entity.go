package parking

import (
	"gotickets/internal/domain/parking/dto"
	"gotickets/internal/domain/parking/types"

	"gorm.io/gorm"
)



type Parking struct {
	gorm.Model
	Name string `json:"name" validate:"required"`
	Type types.ParkingType `json:"type" validate:"required"`
	Total_capacity int `json:"total_capacity" validate:"required"` 
	Price_per_hour int `json:"price_per_hour" validate:"required"`
	Total_reservations int `json:"total_reservations"`
}

func(p *Parking) ToCreateResponse()*dto.CreateResponse{
	return &dto.CreateResponse{
	   Success: true,
	   Message: "Parking created successfully",
	   Data:dto.CreateData{
	   Id: p.ID,
	   Name:p.Name,
	   Type:p.Type,
	   Total_Capacity: p.Total_capacity,
	   Price_per_hour: p.Price_per_hour,
	   CreatedAt: p.CreatedAt.String(),
	   UpdatedAt: p.UpdatedAt.String(),
	   },
	}
}


func(p *Parking) ToGetResponse()*dto.GetResponse{
	return &dto.GetResponse{
		Success: true,
		Message: "Parking fetched successfully",
		Data:[]dto.GetData{{
				Id: p.ID,
				Name:p.Name,
				Type:p.Type,
				Total_Capacity: p.Total_capacity,
				Available_spots: p.Total_capacity - p.Total_reservations,
				Price_per_hour: p.Price_per_hour,
				CreatedAt: p.CreatedAt.String(),
				UpdatedAt: p.UpdatedAt.String(),
	    },
	},
	}
}

func (p *Parking) ToGetSingleResponse() *dto.GetSingleResponse{
	return &dto.GetSingleResponse{
		Success: true,
		Message: "Parking fetched successfully",
		Data:dto.GetData{
		Id: p.ID,
		Name:p.Name,
		Type:p.Type,
		Total_Capacity: p.Total_capacity,
		Available_spots: p.Total_capacity - p.Total_reservations,
		Price_per_hour: p.Price_per_hour,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	},
}}