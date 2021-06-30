package food

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetAllFood)
	g.POST("", h.CreateFood)
}
