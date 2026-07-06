package server

import (
	"fmt"
	"gotickets/internal/config"
	"gotickets/internal/domain/parking"
	"gotickets/internal/domain/reservations"
	"gotickets/internal/domain/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct{
	validate *validator.Validate
}

func (cv *CustomValidator) Validate(i any)error{
	if err := cv.validate.Struct(i);err!=nil{
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func Start(cfg *config.Config, db *gorm.DB){
	db.AutoMigrate(&user.User{},&reservations.Reservation{},&parking.Parking{})

	e:=echo.New()
	e.Validator = &CustomValidator{validate: validator.New()}
	e.Use(middleware.RequestLogger())

	//routes
	user.RegisterRoutes(e,db,cfg)
	parking.RegisterRoutes(e,db,cfg)
	reservations.RegisterRoutes(e,db,cfg)
    
	port:=fmt.Sprintf(":%s",cfg.Port)
	fmt.Printf("Server running on port %s",port)
	if err:=e.Start(port);err!=nil{
		e.Logger.Error("Failed to start server","error",err)
	}
}
