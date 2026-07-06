package parking

import "gorm.io/gorm"

type repo struct {
	db *gorm.DB
}

type Repository interface{
	CreateParking(parking *Parking)error
	GetAllParkings()([]*Parking,error)
	GetParkingById(Id uint)(*Parking,error)
}

func NewRepository(db *gorm.DB) Repository{
	return &repo{db}
}

func(r *repo) CreateParking(parking *Parking)error{
	return r.db.Create(parking).Error
}
func(r *repo) GetAllParkings()([]*Parking,error){
	var parkings []*Parking

	if err:=r.db.Find(&parkings).Error;err!=nil{
		return nil,err
	}

	return parkings,nil


}
func(r *repo) GetParkingById(Id uint)(*Parking,error){
	var parking Parking
    
	if err:=r.db.First(&parking,Id).Error;err!=nil{
		return nil,err
	}

	return &parking,nil
}