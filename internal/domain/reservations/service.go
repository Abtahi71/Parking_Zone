package reservations

import (
	
	"gotickets/internal/domain/reservations/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service{
	return &service{repo}
}

func (s *service) Reserve(req dto.ReserveReq,user_id uint)(*dto.ReserveRes,error){
     reservation:=Reservation{
		User_id:user_id,
		Zone_id:req.Zone_id,
		License_plate:req.License_plate,
	 }
	 

	 result,err:=s.repo.Reserve(&reservation)
	 if err!=nil{
		return nil,err
	 }

	 return result.ToReserveResponse(),err
}

func(s *service) GetMyReservations(user_id uint)(*dto.GetReserveRes,error){
	reservations,err:=s.repo.GetMyReservations(user_id)

	var responses []dto.GetReserveData

	for _,reservation:=range reservations{
		responses = append(responses, dto.GetReserveData{
			Id:reservation.ID,
			License_plate: reservation.License_plate,
			Status: reservation.Status,
			Zone:dto.ZoneData{
				Id:reservation.Zone_id,
				Name:reservation.Zone.Name,
				Type:reservation.Zone.Type,
			},
			CreatedAt: reservation.CreatedAt.String(),
		})
	}
	GetResponse:=&dto.GetReserveRes{
		Success: true,
		Message: "My reservations retrieved successfully",
		//this is a comment in main branch
		Data:responses,
	}
	return GetResponse,err
}

func (s *service)CancelReservation(reservation_id int, user_id uint)(*dto.CancelResponse,error){
	err:=s.repo.CancelReservation(uint(reservation_id),user_id)

	CancelRes:=&dto.CancelResponse{
		Success: true,
		Message: "Reservation cancelled successfully change 1",
	}
	return CancelRes,err

}
func (s *service)GetAllReservations()([]dto.ReserveRes,error){
	reservations,err:=s.repo.GetAllReservations()
	var responses []dto.ReserveRes
	for _,reservation:=range reservations{
		responses = append(responses, *reservation.ToReserveResponse())
	}
	return responses,err
}




