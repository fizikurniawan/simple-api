package route

import (
	"fmt"
	"net/http"
	"os"

	"simple-api/api/app/account"
	"simple-api/api/app/authentication"
	"simple-api/api/app/food"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	var (
		APP     = os.Getenv("APP")
		VERSION = os.Getenv("VERSION")
	)

	// Index
	g.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	authentication.NewHandler().Route(g.Group("/authentication"))
	account.NewHandler().Route(g.Group("/account"))
	food.NewHandler().Route(g.Group("/food"))
}
