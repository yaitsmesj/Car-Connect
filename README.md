# Car-Connect
A small project containing to microservices to Calculate Patrol Cost based on Fuel lid open and close time.

## 1. api-service
api-sercie expses an API that takes two Query Params.
 1. fuellid - true/false
 2. city - {city name}

This informantion is sent to client-service throght Message Broker to calculate Patrol Cost.

It also runs a Scheduler, which send last state of data(fuellid and city) to client-service every 2 minutes.

## 2. client-service
client-service receives Event Message from api-service.
It fetches Today's patrol price for current city and Calculate Patrol Cost based on the time fuellid was open.

API to fetch patrol price is mocked and is present inside client-service itself.

## How to run
Requirement - Go

Run both services separately with -
go run main.go

Call API
  Example - localhost:8090/event?fuellid=false&city=bangalore
