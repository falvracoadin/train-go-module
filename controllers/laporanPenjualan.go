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

func GetLaporanHarian(e echo.Context) error {
	response := &helpers.JSONResponse{}

	tNow := time.Now()
	id := tNow.UnixNano() / 1e10
	req := system.RequestGetParams(e)
	req["id"], _ = strconv.Atoi(req["id"].(string))

	rabbitMQConnection := rmqauto.RabbitMQConnection{}
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(&rabbitMQConnection, structs.MessagePayload{
		id,
		"laporan-harian",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req}, true)
	
	if data == nil {
		code, _ := strconv.Atoi(cd.(string))
		response.StatusCode = code
		response.Message = message.(string)
		response.Data = data
		return e.JSON(code, response)
	}

	code := int(cd.(float64))
	d := data.(map[string]interface{})
	response.StatusCode = code
	response.Data = d["data"]
	response.Message = fmt.Sprintf("%v", message)

	return e.JSON(code, response)
}

func GetLaporanHarianToko(e echo.Context) error {
	response := &helpers.JSONResponse{}

	tNow := time.Now()
	id := tNow.UnixNano() / 1e10
	req := system.RequestGetParams(e)

	rabbitMQConnection := rmqauto.RabbitMQConnection{}
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(&rabbitMQConnection, structs.MessagePayload{
		id,
		"laporan-harian-toko",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req}, true)
	
	if data == nil {
		code, _ := strconv.Atoi(cd.(string))
		response.StatusCode = code
		response.Message = message.(string)
		response.Data = data
		return e.JSON(code, response)
	}

	code := int(cd.(float64))
	d := data.(map[string]interface{})
	response.StatusCode = code
	response.Data = d["data"]
	response.Message = fmt.Sprintf("%v", message)

	return e.JSON(code, response)
}

func GetDetailLaporan(e echo.Context) error {
	response := helpers.JSONResponse{}

	tNow := time.Now()
	id := tNow.UnixNano() / 1e10
	req := system.RequestGetParams(e)

	rabbitMQConnection := rmqauto.RabbitMQConnection{}
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(&rabbitMQConnection, structs.MessagePayload{
		id,
		"rekap-laporan",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req}, true)
	
	if data == nil {
		code, _ := strconv.Atoi(cd.(string))
		response.StatusCode = code
		response.Message = message.(string)
		response.Data = data
		return e.JSON(code, response)
	}

	code := int(cd.(float64))
	d := data.(map[string]interface{})
	response.StatusCode = code
	response.Data = d["data"]
	response.Message = fmt.Sprintf("%v", message)

	return e.JSON(code, response)
}