package messanger

import (
	"fmt"

	"github.com/yaitsmesj/Car-Connect/client-service/calculation"
	"github.com/yaitsmesj/Car-Connect/client-service/config"
	"github.com/yaitsmesj/Car-Connect/client-service/logger"

	"github.com/streadway/amqp"
)

var ch *amqp.Channel
var conn *amqp.Connection
var q amqp.Queue

// ConnectBroker setup connection to RabbitMQ
func ConnectBroker() {
	config := config.GetConfig()
	conn, err := amqp.Dial(config.URL)
	logger.LogMessage(err, "Could Not Connect to Broker", "Successfully connected to Message Broker")
	// defer conn.Close()

	ch, err = conn.Channel()
	logger.LogMessage(err, "Could Not open a Channel", "Successfully opened a Channel")
	// defer ch.Close()

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

func RecvMessage() {
	// ReceiveMessage Receive message on Queue
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	logger.LogMessage(err, "Failed to Register Receiver", "Registered the Receiver")

	for msg := range msgs {
		// Start Go routine
		fmt.Printf("\nMessage: %s\n", msg.Body)
		go calculation.HandleEvent(msg.Body)
	}
}

func CloseConnection() {
	ch.Close()
	conn.Close()
}
