package controller

import (
	"go-api/config"
	"go-api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateTask(c echo.Context) error {
	b := new(model.Task)
	db := config.DB

	// binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	task := &model.Task{
		Title: b.Title,
		Body: b.Body,
		Done: b.Done,
	}

	if err := db.Create(&task).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Data": b,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateTask(c echo.Context) error {
	id := c.Param("id")
	b := new(model.Task)
	db := config.DB

	// binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existing_task := new(model.Task)

	if err := db.First(&existing_task, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existing_task.Title = b.Title
	existing_task.Body = b.Body
	existing_task.Done = b.Done
	if err := db.Save(&existing_task).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Data": existing_task,
	}

	return c.JSON(http.StatusOK, response)
}

func GetAllTasks(c echo.Context) error {
	db := config.DB

	var tasks []model.Task
	if res := db.Find(&tasks); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Datas": tasks,
	}

	return c.JSON(http.StatusOK, response)
}

func GetTask(c echo.Context) error {
	id := c.Param("id")
	db := config.DB

	var task model.Task

	if res := db.First(&task, id); res.Error != nil {
		data := map[string]interface{} {
			"message": res.Error.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{} {
		"Data": task,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	db := config.DB

	var task model.Task

	if err := db.Delete(&task,id).Error; err != nil {
		data := map[string]interface{} {
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{} {
		"Data": "Task with id " + id + " deleted successfully",
	}

	return c.JSON(http.StatusOK, response)
}