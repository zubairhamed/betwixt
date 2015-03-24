package main

import (
	"github.com/zubairhamed/golwm2m"
    "github.com/zubairhamed/goap"
    "log"
)

func main() {
	client := golwm2m.NewClient()

	client.Dial("udp", "localhost", 5683)

    client.Register("golwm2m", func (msg *goap.Message){
        log.Println(goap.CoapCodeToString(msg.Code))
    })
}

