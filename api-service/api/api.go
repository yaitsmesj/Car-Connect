package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/yaitsmesj/api-service/data"
	"github.com/yaitsmesj/api-service/messanger"
)

func lidEvent(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	fuellid, ok := params["fuellid"]
	if !ok || len(fuellid[0]) < 1 {
		http.Error(w, "Missing fuellid Param", 400)
		return
	}

	city, ok := params["city"]
	if !ok || len(city[0]) < 1 {
		http.Error(w, "Missing City Param", 400)
		return
	}
	lid, err := strconv.ParseBool(fuellid[0])
	if err != nil {
		http.Error(w, "Fuellid value should be true or false", 400)
		return
	}
	// Update current state
	data.Fuellid = lid
	data.City = city[0]
	log.Println("Sending Event Message")
	messanger.SendMessage()
	w.Write([]byte("Event Message Sent"))
}

// Start HTTP Server
func StartServer() {
	http.HandleFunc("/event", lidEvent)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
