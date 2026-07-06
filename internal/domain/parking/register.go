package parking

import (
	"gotickets/internal/auth"
	"gotickets/internal/config"
	middleware "gotickets/internal/middlewares"

	"github.com/labstack/echo/v5"

	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo,db *gorm.DB,cfg *config.Config){
	parkingRepo:=NewRepository(db)
	parkingService:=NewService(parkingRepo)
	parkingHandler:=NewHandler(parkingService)

	api:=e.Group("/api/v1/parkings")
	admin:=e.Group("/api/v1/admin/parkings")

	admin.POST("/zones",parkingHandler.CreateParking,middleware.AuthMiddleware(auth.NewJwtService(cfg.JwtSecret)),middleware.RequireRole("admin"))
	api.GET("/all",parkingHandler.GetAllParkings)
	api.GET("/:id",parkingHandler.GetParkingById)

}