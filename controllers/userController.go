package controllers

import (
	"fmt"
	"go-skeleton-manager-rabbitmq/helpers"
	"go-skeleton-manager-rabbitmq/rmqauto"
	"go-skeleton-manager-rabbitmq/structs"
	"go-skeleton-manager-rabbitmq/system"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

const (
	layoutISO = "2006-01-02"
)

func GetUser(e echo.Context) error {
	log.Println("Starting process get user by id")
	response := new(helpers.JSONResponse)

	// Get Parameter user id
	idUser, _ := strconv.Atoi(e.Param("id"))
	var req structs.Filter
	if err := e.Bind(&req); err != nil {
		log.Println(&req)
		response.Message = "Bad Request"
		return e.JSON(http.StatusBadRequest, response)
	}

	if idUser != 0 {
		req.ID = int64(idUser)
	}

	// Post Message to RabbitMQ
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(
		rabbitMQConnection,
		structs.MessagePayload{
			id,
			"get_users",
			tNow.String(),
			"go-skeleton-manager-venturo",
			nil,
			uuid.New().String(),
			map[string]interface{}{
				"id":     req.ID,
				"filter": req.Filter,
				"limit":  req.Limit,
				"offset": req.Offset,
			}}, true)

	// define response data
	d := data.(map[string]interface{})
	response.Data = d["data"]
	response.Message = fmt.Sprintf("%v", message)
	response.StatusCode = int(cd.(float64))

	return e.JSON(response.StatusCode, response)
}

func CreateUser(e echo.Context) error {
	log.Println("Starting process Add user")
	response := new(helpers.JSONResponse)

	// Get Body Request
	var req structs.User

	if err := e.Bind(&req); err != nil {
		log.Println(&req)
		response.Message = "Bad Request"
		return e.JSON(http.StatusBadRequest, response)
	}

	path := "../img/"
	file, _ := system.UploadBase64ToImg(req.File, path, "user-")

	req.Path = path
	req.Foto = ""
	if file != nil {
		req.Foto = file.(string)
	}

	// Post Message to RabbitMQ
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	birthDate, _ := time.Parse(layoutISO, req.BirthDate)
	log.Println(req)

	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"create_users",
		tNow.String(),
		"go-skeleton-manager-venturo",
		nil,
		uuid.New().String(),
		map[string]interface{}{
			"name":       req.Name,
			"address":    req.Address,
			"birth_date": birthDate,
			"email":      req.Email,
			"path":       req.Path,
			"foto":       req.Foto,
			"username":   req.Username,
			"password":   req.Password,
			"created_at": tNow,
		}}, true)

	// define response data
	d := data.(map[string]interface{})
	response.Data = d["data"]
	response.Message = fmt.Sprintf("%v", message)
	response.StatusCode = int(cd.(float64))

	return e.JSON(response.StatusCode, response)
}

func UpdateUser(e echo.Context) error {
	log.Println("Starting process update user by id")
	response := new(helpers.JSONResponse)

	// Get Parameter user id
	idUser, _ := strconv.Atoi(e.Param("id"))

	var req structs.User
	if err := e.Bind(&req); err != nil {
		log.Println(&req)
		response.Message = "Bad Request"
		return e.JSON(http.StatusBadRequest, response)
	}

	path := "../img/"
	file, _ := system.UploadBase64ToImg(req.File, path, "user-")

	req.Path = path
	req.Foto = ""
	if file != nil {
		req.Foto = file.(string)
	}

	// Post Message to RabbitMQ
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	birthDate, _ := time.Parse(layoutISO, req.BirthDate)
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{id,
		"update_users_by_id",
		tNow.String(),
		"go-skeleton-manager-venturo",
		nil,
		uuid.New().String(),
		map[string]interface{}{
			"name":       req.Name,
			"address":    req.Address,
			"birth_date": birthDate,
			"email":      req.Email,
			"path":       req.Path,
			"foto":       req.Foto,
			"username":   req.Username,
			"password":   req.Password,
			"updated_at": tNow,
			"id":         idUser,
		}}, true)

	// define response data
	d := data.(map[string]interface{})
	response.Data = d["data"]
	response.Message = fmt.Sprintf("%v", message)
	response.StatusCode = int(cd.(float64))

	return e.JSON(response.StatusCode, response)
}

func DeleteUser(e echo.Context) error {
	log.Println("Starting process delete user by id")
	response := new(helpers.JSONResponse)

	// Get Parameter user id
	idUser, _ := strconv.Atoi(e.Param("id"))

	// Post Message to RabbitMQ
	tNow := time.Now()
	id := tNow.UnixNano() / 10000000000
	rabbitMQConnection := new(rmqauto.RabbitMQConnection)
	rabbitMQConnection.GlobalConn()
	cd, message, data := rmqauto.RabbitMQRPC(rabbitMQConnection, structs.MessagePayload{
		id,
		"delete_users_by_id",
		tNow.String(),
		"go-skeleton-manager-venturo",
		nil,
		uuid.New().String(),
		map[string]interface{}{
			"id": idUser,
		}}, true)

	// define response data
	d := data.(map[string]interface{})
	response.Data = d["data"]
	response.Message = fmt.Sprintf("%v", message)
	response.StatusCode = int(cd.(float64))

	return e.JSON(response.StatusCode, response)
}

//func TestExport(e echo.Context) error {
//	log.Println("Starting process Export PDF")
//	//response := new(helpers.JSONResponse)
//
//	var header = map[string]string{
//		//cell : value
//		"B1": "Id",
//		"C1": "Nama",
//	}
//
//	var data = []map[string]interface{}{
//		{
//			"column": "B",
//			"id":     1,
//			"nama":   "Sase",
//		},
//		{
//			"column": "C",
//			"id":     2,
//			"nama":   "Aditya",
//		},
//	}
//
//	file := system.ExportExcel(e, "Sheet One", header, nil, data)
//
//	return file
//}
