package user

import (
	"gotickets/internal/auth"
	"gotickets/internal/config"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config){
	userRepo:=NewRepository(db)
	jwtService:=auth.NewJwtService(cfg.JwtSecret)
	userService:=NewService(userRepo,jwtService)
	userHandler:=NewUserHandler(userService)

	api:=e.Group("/api/v1/users")

	api.POST("/register",userHandler.CreateUser)
	api.POST("/login",userHandler.Login)

}