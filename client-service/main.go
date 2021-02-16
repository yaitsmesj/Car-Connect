package main

import (
	"github.com/yaitsmesj/Car-Connect/client-service/api"
	"github.com/yaitsmesj/Car-Connect/client-service/messanger"
)

func main() {

	// Starts Broker
	messanger.ConnectBroker()
	defer messanger.CloseConnection()

	// Start Price API server
	go api.StartPriceServer()

	// Start Receiving Messages
	messanger.RecvMessage()

}
