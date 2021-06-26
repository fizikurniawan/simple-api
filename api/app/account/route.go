package account

import (
	mddlwr "simple-api/api/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("/profile", h.GetProfile, mddlwr.Authorization)
	g.PUT("/profile", h.UpdateProfile, mddlwr.Authorization)
}
