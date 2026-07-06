package parking

import "gotickets/internal/domain/parking/dto"

type service struct {
	repo Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) CreateParking(req dto.CreateParking)(*dto.CreateResponse,error) {
	parking:=Parking{
		Name:req.Name,
		Type:req.Type,
		Total_capacity: req.Total_Capacity,
		Price_per_hour: req.Price_per_hour,
	}

	if err:=s.repo.CreateParking(&parking);err!=nil{
		return nil,err
	}

	return parking.ToCreateResponse(),nil
}

func (s *service) GetAllParkings()(*dto.GetResponse,error){
	parkings,err:=s.repo.GetAllParkings()

	if err!=nil{
		return nil,err
	}

	var responses []dto.GetData

	for _,parking:=range parkings{
		responses = append(responses,dto.GetData{
			Id:parking.ID,
			Name:parking.Name,
			Type:parking.Type,
			Total_Capacity:parking.Total_capacity,
			Available_spots:parking.Total_capacity - parking.Total_reservations,
			Price_per_hour:parking.Price_per_hour,
			CreatedAt:parking.CreatedAt.String(),
			
		})
	}

	GetRes:=&dto.GetResponse{
		Success: true,
		Message: "All parkings retrieved successfully",
		Data:responses,
	}

	return GetRes,nil
}

func (s *service) GetParkingById(Id uint)(*dto.GetSingleResponse,error){
	parking,err:=s.repo.GetParkingById(Id)

	if err!=nil{
		return nil,err
	}

	return parking.ToGetSingleResponse(),nil
}
