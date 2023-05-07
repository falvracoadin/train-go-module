package controllers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"go-skeleton-manager-rabbitmq/helpers"
	"go-skeleton-manager-rabbitmq/rmqauto"
	"go-skeleton-manager-rabbitmq/structs"
	"go-skeleton-manager-rabbitmq/system"
	"log"
	"strconv"
	"time"
)

func LapProdukPerKota(e echo.Context) error {
	log.Println("Starting process laporan jumlah produk, jumlah toko per kota")
	response := new(helpers.JSONResponse)
	// Post Message to RabbitMQ
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	//date, _ := time.Parse(layoutISO, req.CreatedAt)
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	req := system.RequestGetParams(e)

	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"produk-perkota",
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
func LapProdukPerToko(e echo.Context) error {
	log.Println("Starting process laporan jumlah produk per toko per kategori per kota")
	response := new(helpers.JSONResponse)
	// Post Message to RabbitMQ
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	//date, _ := time.Parse(layoutISO, req.CreatedAt)
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	req := system.RequestGetParams(e)

	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"produk-pertoko",
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
func LapRekapPenjualan(e echo.Context) error {
	log.Println("Starting process laporan rekap penjualan")
	response := new(helpers.JSONResponse)
	// Post Message to RabbitMQ
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	//date, _ := time.Parse(layoutISO, req.CreatedAt)
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	req := system.RequestGetParams(e)

	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"rekap-penjualan",
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
