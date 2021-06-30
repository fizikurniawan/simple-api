package food

import (
	"simple-api/api/model"
	res "simple-api/libs/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler() *handler {
	service := NewService()
	return &handler{service}
}

func (h *handler) GetAllFood(c echo.Context) (err error) {
	pagination := new(model.PaginationDTO)

	if err = c.Bind(pagination); err != nil {
		return res.RespError(c, &res.ErrBadRequest)
	}
	if err = c.Validate(pagination); err != nil {
		return res.RespError(c, &res.ErrValidation)
	}

	foods, err := h.service.GetFoods(pagination)
	if err != nil {
		return res.RespError(c, err)
	}
	return res.RespSuccess(c, "Get Food success", foods)
}

func (h *handler) CreateFood(c echo.Context) (err error) {
	dto := new(FoodDTO)
	if err = c.Bind(dto); err != nil {
		return res.RespError(c, &res.ErrBadRequest)
	}

	if err = c.Validate(dto); err != nil {
		return res.RespError(c, &res.ErrValidation)
	}
	food, err := h.service.CreateFood(dto)

	if err != nil {
		return res.RespError(c, err)
	}

	return res.RespSuccess(c, "Success create food", food)
}
