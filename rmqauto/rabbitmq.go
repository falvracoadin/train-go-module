package rmqauto

import "os"

type RabbitMQConnection struct {
	Host, Port, Username, Password, VirtualHost, QueueName string
}

type RabbitMQDefaultPayload struct {
	Route string      `json:"command"`
	Param interface{} `json:"param"`
	Data  interface{} `json:"data"`
}

func (connection *RabbitMQConnection) GlobalConn() {
	connection.Host = os.Getenv("RQ_HOST")
	connection.Port = os.Getenv("RQ_PORT")
	connection.Username = os.Getenv("RQ_USERNAME")
	connection.Password = os.Getenv("RQ_PASSWORD")
	connection.VirtualHost = os.Getenv("RQ_VHOST")
	connection.QueueName = os.Getenv("RQ_QUEUE")
}