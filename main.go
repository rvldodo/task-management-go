package main

import (
	"go-api/config"
	"go-api/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func (c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	// connect to database
	config.DatabaseInit()
	gorm := config.DB

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	// connect routes
	routes.TaskRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}