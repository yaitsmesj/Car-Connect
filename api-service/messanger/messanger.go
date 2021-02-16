package messanger

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/yaitsmesj/Car-Connect/api-service/config"
	"github.com/yaitsmesj/Car-Connect/api-service/data"
	"github.com/yaitsmesj/Car-Connect/api-service/logger"
)

var ch *amqp.Channel
var conn *amqp.Connection
var q amqp.Queue

// ConnectBroker setup connection to RabbitMQ
func ConnectBroker() {
	config := config.GetConfig()
	conn, err := amqp.Dial(config.URL)
	logger.LogMessage(err, "Could Not Connect to Broker", "Successfully connected to Message Broker")

	ch, err = conn.Channel()
	logger.LogMessage(err, "Could Not open a Channel", "Successfully opened a Channel")

	q, err = ch.QueueDeclare(
		"ConnectQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	logger.LogMessage(err, "Failed to declare the Queue", "Declared the queue")
	return
}

// SendMessage publish message on Queue
func SendMessage() {

	msg := fmt.Sprintf(`{"Fuellid":"%t", "City":"%s"}`, data.Fuellid, data.City)
	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	logger.LogMessage(err, "Failed to Publish message", "Published the message")
}

func CloseConnection() {
	ch.Close()
	conn.Close()
}
