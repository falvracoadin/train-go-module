package controllers

import (
	"fmt"
	"go-skeleton-manager-rabbitmq/helpers"
	"go-skeleton-manager-rabbitmq/rmqauto"
	"go-skeleton-manager-rabbitmq/structs"
	"go-skeleton-manager-rabbitmq/system"
	"strconv"
	"time"
	"log"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func CreatePenjualan(e echo.Context) error {
	response := new(helpers.JSONResponse)

	req := system.RequestGetParams(e)
	now := time.Now()
	id := now.UnixNano() / 10000000000
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"create_penjualan",
		now.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req},
		true)

	if data == nil {
		code, _ := strconv.Atoi(cd.(string))
		response.Data = data
		response.Message = message.(string)
		response.StatusCode = code
		return e.JSON(response.StatusCode, response)
	}

	d := data.(map[string]interface{})
	response.Data = d["data"]
	response.Message = fmt.Sprintf("%v", message)
	response.StatusCode = int(cd.(float64))

	return e.JSON(response.StatusCode, response.Data)
}

func UpdatePenjualan( e echo.Context) error {
	response := new(helpers.JSONResponse)

	IdData, _ := strconv.Atoi(e.Param("id"))
	req := system.RequestGetParams(e)
	req["id"] = IdData
	log.Println(req)

	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"update_penjualan",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req},
		true)
	
	if data == nil{
		code, _ := strconv.Atoi(cd.(string))
		response.Data = data
		response.Message = message.(string)
		response.StatusCode = code
		return e.JSON(response.StatusCode, response)
	}

	d := data.(map[string]interface{})
	response.Data = d
	response.Message = fmt.Sprintf("%v", message)
	response.StatusCode = int(cd.(float64))

	return e.JSON(response.StatusCode, response.Data)
}

func DeletePenjualan(e echo.Context) error{
	response:= new(helpers.JSONResponse)

	idDb, _ := strconv.Atoi(e.Param("id"))

	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000

	rabbitMQCOnnection := new(rmqauto.RabbitMQConnection)
	rabbitMQCOnnection.GlobalConn()
	req := system.RequestGetParams(e)
	req["id"] = idDb

	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQCOnnection, structs.MessagePayload{
		id,
		"delete_penjualan",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req}, true)

	if data == nil{
		code, _ := strconv.Atoi(cd.(string))
		response.Data = data
		response.Message = message.(string)
		response.StatusCode = code
		return e.JSON(code, response)
	}

	d := data.(map[string]interface{})
	response.Data = d["data"]
	response.Message = fmt.Sprintf("%v", message)
	response.StatusCode = int(cd.(float64))
	return e.JSON(response.StatusCode, response)
}

func GetListPenjualan(e echo.Context) error {
	response := helpers.JSONResponse{}

	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000

	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	req := system.RequestGetParams(e)
	if req["limit"] != nil{
		limit, _ := strconv.ParseInt(req["limit"].(string), 0, 32)
		req["limit"] = limit
	}
	if req["offset"] != nil {
		offset, _ := strconv.ParseInt(req["offset"].(string), 0, 32)
		req["offset"] = offset
	}

	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"get_list_penjualan",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req}, true)
	
	if data == nil{
		code, _ := strconv.Atoi(cd.(string))
		response.StatusCode = code
		response.Message = message.(string)
		response.Data = data
		return e.JSON(code, response)
	}

	d := data.(map[string]interface{})
	response.Data = d["data"]
	response.Message = fmt.Sprintf("%v", message)
	response.StatusCode = int(cd.(float64))
	return e.JSON(response.StatusCode, response)
}

func GetDetailPenjualan(e echo.Context) error{
	response := new(helpers.JSONResponse)

	req := system.RequestGetParams(e)
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	if req["limit"] != nil{
		limit, _ := strconv.ParseInt(req["limit"].(string), 0, 32)
		req["limit"] = limit
	}
	if req["offset"] != nil{
		offset, _ := strconv.ParseInt(req["offset"].(string), 0, 32)
		req["offset"] = offset
	}
	
	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"get_detail_penjualan_detail",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil, 
		uuid.New().String(),
		req}, true)
	
	if data == nil{
		code, _ := strconv.Atoi(cd.(string))
		response.StatusCode = code
		response.Message = message.(string)
		response.Data = data
		return e.JSON(code, response)
	}
	d := data.(map[string]interface{})
	response.StatusCode = int(cd.(float64))
	response.Data = d["data"]
	response.Message = message.(string)
	return e.JSON(response.StatusCode, response)
}