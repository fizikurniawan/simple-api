package authentication

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

func (h *handler) Register(c echo.Context) (err error) {
	dto := new(RegisterDTO)

	if err = c.Bind(dto); err != nil {
		return res.RespError(c, &res.ErrBadRequest)
	}
	if err = c.Validate(dto); err != nil {
		fmt.Println(err)
		return res.RespError(c, &res.ErrValidation)
	}

	user, err := h.service.Register(dto)
	if err != nil {
		return res.RespError(c, err)
	}

	return res.RespSuccess(c, "Register success", user)
}

func (h *handler) Login(c echo.Context) (err error) {
	dto := new(LoginDTO)

	if err = c.Bind(dto); err != nil {
		return res.RespError(c, &res.ErrBadRequest)
	}

	if err = c.Validate(dto); err != nil {
		return res.RespError(c, &res.ErrValidation)
	}

	data, err := h.service.Login(dto)
	if err != nil {
		return res.RespError(c, err)
	}

	return res.RespSuccess(c, "Login success", data)

}
