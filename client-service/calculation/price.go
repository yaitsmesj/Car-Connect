package calculation

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

type Data struct {
	Fuellid bool
	City    string
}

var stime, etime time.Time
var fuellid bool
var cprice float32
var price float32

// Calculate Price
func HandleEvent(msg []byte) {
	var data Data
	err := unmarshal(msg, &data)
	if err != nil {
		log.Printf("Could not convert message to struct: %v\n", err)
		return
	}

	if data.Fuellid {
		if fuellid {
			etime = time.Now()
			//calculate price
			calPrice()
		} else {
			fuellid = true
			stime = time.Now()
			//Fetch price
			price = GetPrice(data.City)
			fmt.Printf("Current Patrol Price: %.2f\n", price)
		}
	} else {
		if fuellid {
			fuellid = false
			etime = time.Now()
			//Calculate price
			calPrice()
			stime = time.Time{}
			fmt.Printf("Total Cost : %.2f\n", cprice)
		}
	}
}

func calPrice() {
	d := etime.Sub(stime)
	fd := math.Ceil(d.Seconds())
	q := fd / 30
	cprice = price * float32(q)
	fmt.Printf("Time: %v Seconds, Oil quantity: %.2f Litres, Cost: %.2f\n", fd, q, cprice)
}

func unmarshal(msg []byte, data *Data) error {
	d := struct {
		Fuellid string
		City    string
	}{}
	err := json.Unmarshal(msg, &d)
	if err != nil {
		return err
	}
	data.City = d.City
	data.Fuellid, err = strconv.ParseBool(d.Fuellid)
	return err
}
