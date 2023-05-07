package rmqauto

import (
	"encoding/json"
	"go-skeleton-manager-rabbitmq/helpers"
	"log"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

func RabbitMQRPC(connection *RabbitMQConnection, body interface{}, fReply bool) (responseStatus interface{}, responseMessage interface{}, responseData interface{}) {
	url := connection.Host + ":" + connection.Port + connection.VirtualHost

	log.Println("[AMQP] " + url + " | " + connection.QueueName)

	dial, err := amqp.Dial("amqp://" + connection.Username + ":" + connection.Password + "@" + url)
	if err != nil {
		helpers.HandleError("failed to connect rabbitmq", err)
	}
	defer dial.Close()

	channel, err := dial.Channel()
	if err != nil {
		helpers.HandleError("failed to open a channel in rabbitmq", err)
	}
	defer channel.Close()

	corrId := helpers.RandomByte(8)
	quenameRd := "go-skeleton/" + strconv.Itoa(int(time.Now().Unix())) + corrId

	queue, err := channel.QueueDeclare(
		quenameRd,
		true,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		helpers.HandleError("failed to declare a queue in rabbitmq", err)
	}
	defer channel.QueueDelete(quenameRd, false, false, true)

	message, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		helpers.HandleError("failed to register a consumer in rabbitmq", err)
	}

	_, err = channel.QueueDeclare(
		connection.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		helpers.HandleError("failed to declare a queue in rabbitmq", err)
	}

	waitResponse := true
	if fReply == false {
		waitResponse = false
		queue.Name = ""
	}

	err = channel.Publish(
		"",
		connection.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       queue.Name,
			Body:          []byte(helpers.JSONEncode(body)),
			Expiration:    "60000",
		},
	)
	log.Println("publish :", helpers.JSONEncode(body))
	if err != nil {
		helpers.HandleError("failed to publish a message to rabbitmq", err)
	}

	rabbitTimeout := time.After(10 * time.Second)

	for waitResponse {
		select {
		case <-rabbitTimeout:
			responseStatus = "500"
			responseMessage = "RPC timeout " + connection.QueueName
			waitResponse = false
		case data := <-message:
			log.Println("reply : ", string(data.Body))
			log.Println("")

			if corrId == data.CorrelationId {
				if helpers.JSONEncode(body) == string(data.Body) {
					responseStatus = "error"
					responseMessage = "the rpc server did not respond"
					responseData = nil

				} else {
					var response map[string]interface{}
					json.Unmarshal([]byte(string(data.Body)), &response)
					responseStatus = response["Status"]
					responseMessage = response["StatusMessage"]
					responseData = response
				}
				waitResponse = false
			}
		}
	}
	return
}
