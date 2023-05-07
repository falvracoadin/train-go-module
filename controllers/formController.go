package controllers

import(
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"go-skeleton-manager-rabbitmq/helpers"
	"go-skeleton-manager-rabbitmq/rmqauto"
	"go-skeleton-manager-rabbitmq/structs"
	"go-skeleton-manager-rabbitmq/system"
	"strconv"
	"time"
)

func GetForm(e echo.Context) error {
	response := new(helpers.JSONResponse)

	idData, _ := strconv.Atoi(e.Param("id"))
	req := system.RequestGetParams(e)
	req["id"] = idData

	tNow := time.Now()
	id := tNow.UnixNano() /1e10
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"get-form",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req}, true)
	
	if data == nil{
		code, _ :=strconv.Atoi(cd.(string))
		response.Data = data
		response.Message = message.(string)
		response.StatusCode = code
		return e.JSON(code, response)
	}

	d := data.(map[string]interface{})
	response.Data = d["data"]
	response.Message = fmt.Sprintf("%v", message)
	response.StatusCode = int(cd.(float64))
	return e.JSON(response.StatusCode, response.Data)

}