package reservations

import (
	"gotickets/internal/auth"
	"gotickets/internal/config"
	middleware "gotickets/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo,db *gorm.DB,cfg *config.Config) {
	repo:=NewRepository(db)
	jwtService:=auth.NewJwtService(cfg.JwtSecret)
	service:=NewService(repo)

	handler:=NewHandler(service)
	admin:=e.Group("/api/v1/admin")
	admin.Use(middleware.AuthMiddleware(jwtService),middleware.RequireRole("admin"))

	api:=e.Group("/api/v1/reservations")

	api.POST("/reserve",handler.Reserve,middleware.AuthMiddleware(jwtService))
	api.GET("/myreservations",handler.GetMyReservations,middleware.AuthMiddleware(jwtService))
	api.GET("/allreservations",handler.GetAllReservations)
	api.DELETE("/cancel/:id",handler.CancelReservation,middleware.AuthMiddleware(jwtService))
	

}