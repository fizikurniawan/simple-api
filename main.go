package main

import (
	"os"
	customMiddleware "simple-api/api/middleware"
	"simple-api/api/route"
	db "simple-api/database"
	"simple-api/libs/util"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	godotenv.Load(".env")
}

func main() {

	var (
		PORT = os.Getenv("PORT")
	)

	// Init
	db.Init()
	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = customMiddleware.ErrorHandler

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	// Routes
	route.Init(e.Group(""))

	// Start
	e.Logger.Fatal(e.Start(":" + PORT))

}
