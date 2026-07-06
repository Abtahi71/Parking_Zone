package parking

import (
	"gotickets/internal/config"

	"github.com/labstack/echo/v5"

	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo,db *gorm.DB,cfg *config.Config){
	parkingRepo:=NewRepository(db)
	parkingService:=NewService(parkingRepo)
	parkingHandler:=NewHandler(parkingService)

	api:=e.Group("/api/v1/parkings")

	api.POST("/create",parkingHandler.CreateParking)
	api.GET("/all",parkingHandler.GetAllParkings)
	api.GET("/:id",parkingHandler.GetParkingById)

}