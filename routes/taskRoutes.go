package routes

import (
	"go-api/controller"

	"github.com/labstack/echo/v4"
)

func TaskRoutes(e *echo.Echo) {

	taskRoute := e.Group("/task")
	taskRoute.GET("/", controller.GetAllTasks)
	taskRoute.GET("/:id", controller.GetTask)
	taskRoute.POST("/", controller.CreateTask)
	taskRoute.PATCH("/:id", controller.UpdateTask)
	taskRoute.DELETE("/:id", controller.DeleteTask)

}