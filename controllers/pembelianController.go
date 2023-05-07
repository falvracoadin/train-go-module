package controllers

import (
	"fmt"
	"go-skeleton-manager-rabbitmq/helpers"
	"go-skeleton-manager-rabbitmq/rmqauto"
	"go-skeleton-manager-rabbitmq/structs"
	"go-skeleton-manager-rabbitmq/system"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func CreatePembelian(e echo.Context) error {
	log.Println("Starting process Create Pembelian")
	response := new(helpers.JSONResponse)

	req := system.RequestGetParams(e)
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"create_pembelian",
		tNow.String(),
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

	return e.JSON(response.StatusCode, response)
}

func UpdatePembelian(e echo.Context) error {
	log.Println("Starting process Update Pembelian")
	response := new(helpers.JSONResponse)

	// Get Id Data
	IdDb, _ := strconv.Atoi(e.Param("id"))
	req := system.RequestGetParams(e)
	req["id"] = IdDb

	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"update_pembelian",
		tNow.String(),
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

	return e.JSON(response.StatusCode, response)
}

func DeletePembelian(e echo.Context) error {
	log.Println("Starting process Delete Pembelian")
	response := new(helpers.JSONResponse)

	idDb, _ := strconv.Atoi(e.Param("id"))

	// Post Message to RabbitMQ
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	//date, _ := time.Parse(layoutISO, req.CreatedAt)
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	req := system.RequestGetParams(e)
	req["id"] = idDb

	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"delete_pembelian",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req}, true)

	// define response data
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

	return e.JSON(response.StatusCode, response)
}

func GetListPembelianById(e echo.Context) error {
	log.Println("Starting process Get Detail Pembelian")
	response := new(helpers.JSONResponse)

	idDb := e.Param("id")

	// Post Message to RabbitMQ
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	//date, _ := time.Parse(layoutISO, req.CreatedAt)
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	req := system.RequestGetParams(e)
	req["id"], _ = strconv.Atoi(idDb)

	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"get_list_pembelian_by_id",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req}, true)

	// define response data
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

	return e.JSON(response.StatusCode, response)
}

func GetDetailPembelianDetail(e echo.Context) error {
	log.Println("Starting process Get Detail Pembelian Detail")
	response := new(helpers.JSONResponse)

	idDb := e.Param("id")

	// Post Message to RabbitMQ
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	//date, _ := time.Parse(layoutISO, req.CreatedAt)
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	req := system.RequestGetParams(e)
	req["id"], _ = strconv.Atoi(idDb)

	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"get_detail_pembelian_detail",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req}, true)

	// define response data
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

	return e.JSON(response.StatusCode, response)
}

func GetListPembelian(e echo.Context) error {
	log.Println("Starting process Get List Pembelian")
	response := new(helpers.JSONResponse)

	// Post Message to RabbitMQ
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	//date, _ := time.Parse(layoutISO, req.CreatedAt)
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	req := system.RequestGetParams(e)
	if req["limit"] != nil {
		limit, _ := strconv.ParseInt(req["limit"].(string), 0, 32)
		req["limit"] = int32(limit)
	}
	if req["offset"] != nil {
		offset, _ := strconv.ParseInt(req["offset"].(string), 0, 32)
		req["offset"] = int32(offset)
	}

	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"get_list_pembelian",
		tNow.String(),
		"go-skeleton-manager-rabbitmq",
		nil,
		uuid.New().String(),
		req}, true)

	// define response data
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
	return e.JSON(response.StatusCode, response)
}

func GetPembelianWithDetail(e echo.Context) error {
	log.Println("Starting process Get List Pembelian With Detail")
	response := new(helpers.JSONResponse)

	idDb := e.Param("id")
	req := system.RequestGetParams(e)
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	if req["limit"] != nil {
		limit, _ := strconv.ParseInt(req["limit"].(string), 0, 32)
		req["limit"] = int32(limit)
	}
	if req["offset"] != nil {
		offset, _ := strconv.ParseInt(req["offset"].(string), 0, 32)
		req["offset"] = int32(offset)
	}
	req["id"], _ = strconv.Atoi(idDb)

	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"get_pembelian_with_detail",
		tNow.String(),
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

	return e.JSON(response.StatusCode, response)
}
