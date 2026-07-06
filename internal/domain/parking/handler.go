package parking

import (
	"fmt"
	"gotickets/internal/domain/parking/dto"
	ErrorResponse "gotickets/internal/httpresponse"
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

func (h *handler) CreateParking(c *echo.Context)error{
     var req dto.CreateParking
	 if err:=c.Bind(&req);err!=nil{
		return c.JSON(http.StatusBadRequest,ErrorResponse.Error{
			Code: http.StatusBadRequest,
			Message: "Invalid request body",
		})
	 }
	 if err:=c.Validate(&req);err!=nil{
		return c.JSON(http.StatusBadRequest,ErrorResponse.Error{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	 }
	 response,err:=h.service.CreateParking(req)
	 if err!=nil{
		return c.JSON(http.StatusInternalServerError,ErrorResponse.Error{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	 }
     fmt.Printf("%+v\n", response)
	 return c.JSON(http.StatusCreated,response)
}

func (h *handler) GetAllParkings(c *echo.Context)error{
	parkings,err:=h.service.GetAllParkings()

	if err!=nil{
		return err
	}

	return c.JSON(http.StatusOK,parkings)
}
func (h *handler) GetParkingById(c *echo.Context)error{
	id,err:=strconv.Atoi(c.Param("id"))
	

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid event id",
			Details: err.Error(),
		})
	}

	parking,err:=h.service.GetParkingById(uint(id))

	if err!=nil{
		return err
	}

	return c.JSON(http.StatusOK,parking)
}