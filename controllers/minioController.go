package controllers

import (
	// "fmt"
	// "github.com/google/uuid"
	"go-skeleton-manager-rabbitmq/miniogo"

	"github.com/labstack/echo"

	// "go-skeleton-manager-rabbitmq/rmqauto"
	// "go-skeleton-manager-rabbitmq/structs"
	"go-skeleton-manager-rabbitmq/helpers"
	// "go-skeleton-manager-rabbitmq/system"
	// "strconv"
	// "time"
)

func UploadFile(e echo.Context) error {
	response := new(helpers.JSONResponse)
	response.StatusCode = 200
	response.Message = "test"
	file, _ := e.FormFile("file")
	res, _ := miniogo.Upload(file)
	response.Data = res
	return e.JSON(200, response)
}