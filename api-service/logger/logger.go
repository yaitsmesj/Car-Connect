package logger

import (
	"fmt"
	"log"
)

func LogMessage(err error, errmsg string, sucmsg string) {
	if err != nil {
		log.Fatalf("%s: %s", errmsg, err)
	} else {
		fmt.Printf("%s\n", sucmsg)
	}
}
