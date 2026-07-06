package reservations

import (
	"fmt"
	"gotickets/internal/domain/parking"
	"gotickets/internal/domain/reservations/types"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repo struct {
	db *gorm.DB
}

type Repository interface{
	Reserve(reservation *Reservation)(*Reservation,error)
	GetMyReservations(user_id uint)([]*Reservation,error)
	CancelReservation(reservation_id uint,user_id uint)error
	GetAllReservations() ([]*Reservation,error)
}

func NewRepository(db *gorm.DB) Repository{
	return &repo{db}
}

func(r *repo) Reserve(reservation *Reservation) (*Reservation,error){
   var reservations Reservation

   err:=r.db.Transaction(func(tx *gorm.DB)error{
	var parking parking.Parking

	err:=tx.Clauses(clause.Locking{Strength:"Update"}).First(&parking,reservation.Zone_id).Error
	if err!=nil{
		return err
	}
	if parking.Total_reservations>=parking.Total_capacity{
		return fmt.Errorf("No more space available in this spot")
	}
	reservations=Reservation{
		User_id:reservation.User_id,
		Zone_id:reservation.Zone_id,
		License_plate:reservation.License_plate,
		Status:types.Active,
	}

	err=tx.Create(&reservations).Error
	if err!=nil{
		return err
	}
	parking.Total_reservations+=1
	if err:=tx.Save(&parking).Error;err!=nil{
		return err
	}
	return nil
   })

	if err!=nil{
		return nil,err
	}
	return &reservations,nil

}

func (r *repo) GetMyReservations(user_id uint)([]*Reservation,error){
     var reservations []*Reservation

	 err:=r.db.Preload("Zone").Where("User_id=?",user_id).Find(&reservations).Error
	 if err!=nil{
		return nil,err
	 }
	 return reservations,nil
}
func (r *repo) CancelReservation(reservation_id uint,user_id uint)error{
     

	 err:=r.db.Transaction(func(tx *gorm.DB)error{
		var reservation Reservation
		err:=tx.Clauses(clause.Locking{Strength: "Update"}).
		Where("Id=? AND User_id=?",reservation_id,user_id).
		First(&reservation).Error
		if err!=nil{
			return err
		}

		var parking parking.Parking

		err=tx.Clauses(clause.Locking{Strength: "Update"}).First(&parking,reservation.Zone_id).Error
		if err!=nil{
			return err
		}

		parking.Total_reservations-=1
		if err:=tx.Save(&parking).Error;err!=nil{
			return err
		}
		reservation.Status=types.Cancelled
		if err:=tx.Save(&reservation).Error;err!=nil{
			return err
		}
		return nil

	 })

	 return err
	 
}

func (r *repo) GetAllReservations() ([]*Reservation,error){
	 var reservations []*Reservation

	 err:=r.db.Preload("Zone,User").Find(&reservations).Error
	 if err!=nil{
		return nil,err
	 }

	 return reservations,nil
	}

