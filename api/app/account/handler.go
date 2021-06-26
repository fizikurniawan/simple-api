package account

import (
	"fmt"
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

func (h *handler) GetProfile(c echo.Context) (err error) {
	email := fmt.Sprint(c.Get("email"))
	data, err := h.service.GetProfile(email)

	if err != nil {
		return res.RespError(c, &res.ErrBadRequest)
	}

	return res.RespSuccess(c, "Success get profile detail", data)
}

func (h *handler) UpdateProfile(c echo.Context) (err error) {
	dto := new(UpdateProfileDTO)
	email := fmt.Sprint(c.Get("email"))

	if err = c.Bind(dto); err != nil {
		return res.RespError(c, &res.ErrBadRequest)
	}

	if err = c.Validate(dto); err != nil {
		return res.RespError(c, &res.ErrValidation)
	}

	_, err = h.service.UpdateProfile(dto, email)
	if err != nil {
		return res.RespError(c, &res.ErrBadRequest)
	}

	if err != nil {
		return res.RespError(c, err)
	}

	return res.RespSuccess(c, "Success update profile", dto)
}
