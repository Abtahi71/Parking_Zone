package user

import (
	"gotickets/internal/domain/user/dto"
	ErrorResponse "gotickets/internal/httresponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	service *service
}

func NewUserHandler(service *service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) CreateUser(c *echo.Context)error{
	var req dto.RegisterRequest
	if err:=c.Bind(&req);err!=nil{
		return c.JSON(http.StatusBadRequest,ErrorResponse.Error{Code: http.StatusBadRequest,Message: "Invalid request body"})
	}
	if err:=c.Validate(&req);err!=nil{
		return c.JSON(http.StatusBadRequest,ErrorResponse.Error{Code: http.StatusBadRequest,Message: err.Error()})
	}
	

	response,err:= u.service.CreateUser(req)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,ErrorResponse.Error{Code: http.StatusInternalServerError,Message: err.Error()})
	}
	return c.JSON(http.StatusCreated,response)
}

func (u *UserHandler) Login(c *echo.Context) error{
	var req dto.LoginRequest

	if err:=c.Bind(&req);err!=nil{
		return c.JSON(http.StatusBadRequest,ErrorResponse.Error{Code: http.StatusBadRequest,Message: "Invalid request body"})
	}
	if err:=c.Validate(&req);err!=nil{
		return c.JSON(http.StatusBadRequest,ErrorResponse.Error{Code: http.StatusBadRequest,Message: err.Error()})
	}
	response,err:= u.service.Login(req)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,ErrorResponse.Error{Code: http.StatusInternalServerError,Message: err.Error()})
	}
	return c.JSON(http.StatusOK,response)
}
