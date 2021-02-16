package api

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func StartPriceServer() {
	http.HandleFunc("/price", sendPrice)
	log.Fatal(http.ListenAndServe(":8091", nil))
}

func sendPrice(w http.ResponseWriter, req *http.Request) {
	params, ok := req.URL.Query()["city"]
	if !ok || len(params[0]) < 1 {
		http.Error(w, "Missing City Param", 400)
		return
	}
	cprice := ((rand.Float32() * 10) + 80)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cprice)
}
