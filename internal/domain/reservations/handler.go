package reservations

import (
	"fmt"
	"gotickets/internal/domain/reservations/dto"
	ErrorResponse "gotickets/internal/httresponse"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service}
}

func (h *handler) Reserve(c *echo.Context) error {
	var req dto.ReserveReq
	if err:=c.Bind(&req);err!=nil{
		return c.JSON(http.StatusBadRequest,ErrorResponse.Error{Code: http.StatusBadRequest,Message: "Invalid request body"})
	}
	if err:=c.Validate(&req);err!=nil{
		return c.JSON(http.StatusBadRequest,ErrorResponse.Error{Code: http.StatusBadRequest,Message: err.Error()})
	}
	user_id:=c.Get("userId").(uint)
	response,err:=h.service.Reserve(req,user_id)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,ErrorResponse.Error{Code: http.StatusInternalServerError,Message: err.Error()})
	}
	return c.JSON(http.StatusCreated,response)
}

func (h *handler) GetMyReservations(c *echo.Context)error{
	user_id:=c.Get("userId").(uint)
	reservations,err:=h.service.GetMyReservations(user_id)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,ErrorResponse.Error{Code: http.StatusInternalServerError,Message: err.Error()})
	}
	return c.JSON(http.StatusOK,reservations)
}

func (h *handler) CancelReservation(c *echo.Context) error{
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		return c.JSON(http.StatusBadRequest,ErrorResponse.Error{Code: http.StatusBadRequest,Message: "Invalid reservation id"})	
	}
	user_id:=c.Get("userId")

	response,err:=h.service.CancelReservation(id,user_id.(uint))
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,ErrorResponse.Error{Code: http.StatusInternalServerError,Message: err.Error()})
	}
	return c.JSON(http.StatusOK,response)
}

func (h *handler) GetAllReservations(c *echo.Context)error{
	reservations,err:=h.service.GetAllReservations()
	fmt.Println("Reservations hit now")
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,ErrorResponse.Error{Code: http.StatusInternalServerError,Message: err.Error()})
	}
	return c.JSON(http.StatusOK,reservations)
}