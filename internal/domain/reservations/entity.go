package reservations

import (
	"gotickets/internal/domain/parking"
	"gotickets/internal/domain/reservations/dto"
	"gotickets/internal/domain/reservations/types"
	"gotickets/internal/domain/user"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	User_id uint 
	Zone_id uint
	License_plate string
	Status types.ReservationStatus 
	Zone parking.Parking `gorm:"foreignKey:Zone_id"`
	User user.User `gorm:"foreignKey:User_id"`
}

func(r *Reservation) ToReserveResponse() *dto.ReserveRes{
	return &dto.ReserveRes{
	Success : true,
	Message : "Reservation created successfully",
	Data:dto.ReserveData{
	Id : r.ID,
	User_id :r.User_id,
	Zone_id :r.Zone_id,
	License_plate :r.License_plate,
	Status  : r.Status,
	CreatedAt :r.CreatedAt.String(),
	UpdatedAt  :r.UpdatedAt.String(),
	},
}
}
func (r *Reservation) ToGetResponse() *dto.GetReserveRes {
	return &dto.GetReserveRes{
		Success: true,
		Message: "Reservation fetched successfully",
		Data: []dto.GetReserveData{{
			Id:             r.ID,
			License_plate:  r.License_plate,
			Status:         r.Status,
			Zone:           dto.ZoneData{Id:r.Zone_id,Name:r.Zone.Name,Type:r.Zone.Type	},
			CreatedAt:      r.CreatedAt.String(),
			UpdatedAt:      r.UpdatedAt.String(),
		},
	},
	}
}