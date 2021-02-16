package main

import (
	"fmt"

	"github.com/yaitsmesj/api-service/api"
	"github.com/yaitsmesj/api-service/messanger"
	"github.com/yaitsmesj/api-service/scheduler"
)

func main() {

	//start API Service

	// Start Message Broker
	messanger.ConnectBroker()
	defer messanger.CloseConnection()

	//start Scheduler
	done := scheduler.Schedule()

	api.StartServer()

	// time.Sleep(6 * time.Minute)
	done <- true
	fmt.Println("Exiting...")
}
