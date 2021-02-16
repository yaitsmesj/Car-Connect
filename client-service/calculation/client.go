package calculation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const dprice float32 = 80.0

type data struct {
	price float32
	ftime time.Time
}

var pricemap = make(map[string]data)

// GetPrice Returns patrol price of a city
func GetPrice(city string) float32 {
	cprice, ok := pricemap[city]
	if !ok || (time.Now().Sub(cprice.ftime).Hours() >= 24) {
		cprice = data{fetchPrice(city), time.Now()}
		pricemap[city] = cprice
	} else {
		fmt.Printf("Returning Cached Price: %.2f", cprice.price)
	}
	return cprice.price
}

func fetchPrice(city string) float32 {
	resp, err := http.Get("http://localhost:8091/price?city=" + city)
	if err != nil {
		log.Printf("Could not fetch current price Error: %v\nUsing Default Price: 80", err)
		return dprice
	}
	var price float32
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Could not fetch current price Error: %v\nUsing Default Price: 80", err)
		return dprice
	}
	err = json.Unmarshal(body, &price)
	if err != nil {
		log.Printf("Could not fetch current price Error: %v\nUsing Default Price: 80", err)
		return dprice
	}
	return price
}
