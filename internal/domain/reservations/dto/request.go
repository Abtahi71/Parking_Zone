package dto

type ReserveReq struct{
	Zone_id uint `json:"zone_id" validate:"required"`
	License_plate string `json:"license_plate" validate:"required"`
}


